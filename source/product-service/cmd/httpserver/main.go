package main

import (
	"clean-code-structure/config"
	"clean-code-structure/delivery/httpserver"
	"clean-code-structure/logger"
	"clean-code-structure/repository/migrator"
	"clean-code-structure/repository/pgsql"
	"clean-code-structure/repository/pgsql/pgsqlproduct"
	"clean-code-structure/repository/seeder"
	"clean-code-structure/service/healthservice"
	"clean-code-structure/service/productservice"
	"clean-code-structure/service/transactionservice"
	"clean-code-structure/validator/healthvalidator"
	"clean-code-structure/validator/productvalidator"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	done := make(chan bool)

	cfg, err := config.Load("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cfg: %+v\n", cfg)

	logger.Start(cfg.Logger)

	mgr := migrator.New(cfg.PgSQL)
	mgr.Down()
	mgr.Up()

	rSvcs, rVal := setupServices(cfg, wg, done)

	sdr := seeder.New(rSvcs)
	sdr.Seed()

	server := httpserver.New(cfg, rSvcs, rVal)
	go func() {
		server.Serve()
	}()

	if cfg.Application.EnableProfiling {
		profiling(cfg, wg, done)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Quit signal received")
	close(done)

	fmt.Println("Wait for done all processes")
	wg.Wait()

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx,
		time.Duration(cfg.Application.GracefulShutdownTimeoutInSecond)*time.Second)
	defer cancel()

	fmt.Println("Shutting down server router")
	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	fmt.Println("received interrupt signal, shutting down gracefully..")
	<-ctxWithTimeout.Done()
}

func profiling(cfg config.Config, wg *sync.WaitGroup, done <-chan bool) {
	fmt.Printf("Profiling enabled on port %d", cfg.Application.ProfilingPort)
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Application.ProfilingPort),
		ReadHeaderTimeout: time.Duration(cfg.Application.TimeoutSeconds) * time.Second,
	}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	go func() {
		<-done
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Application.TimeoutSeconds)*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
}

func setupServices(cfg config.Config, wg *sync.WaitGroup, done chan bool) (requiredServices httpserver.RequiredServices, requiredValidators httpserver.RequiredValidators) {

	pgsqlRepo := pgsql.New(cfg.PgSQL)

	productRepo := pgsqlproduct.New(pgsqlRepo)

	requiredValidators.HealthValidator = healthvalidator.New()
	requiredValidators.ProductValidator = productvalidator.New(productRepo)

	requiredServices.HealthService = healthservice.New()
	requiredServices.TransactionService = transactionservice.New()
	requiredServices.ProductService = productservice.New(productRepo, requiredServices.TransactionService)

	return requiredServices, requiredValidators
}
