package pgsql

import (
	"clean-code-structure/logger"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Name     string `koanf:"name"`
	SslMode  string `koanf:"ssl_mode"`
}

type PGSQLDB struct {
	config Config
	db     *sql.DB
}

type Transaction interface {
	GetRollbackChannel() <-chan bool
	GetCommitChannel() <-chan bool
	Rollback()
	Commit()
}

func New(config Config) *PGSQLDB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User, config.Password, config.Host, config.Port, config.Name, config.SslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &PGSQLDB{
		config: config,
		db:     db,
	}
}

func (p *PGSQLDB) Conn() *sql.DB {
	return p.db
}

func (p *PGSQLDB) WaitForTransaction(ctx context.Context, transaction Transaction, tx *sql.Tx) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				err := tx.Rollback()
				if err != nil {
					logger.Logger.Error("Error on rollback: " + err.Error())
				}
				return
			case <-transaction.GetRollbackChannel():
				err := tx.Rollback()
				if err != nil {
					logger.Logger.Error("Error on rollback: " + err.Error())
				}
				return
			case <-transaction.GetCommitChannel():
				err := tx.Commit()
				if err != nil {
					logger.Logger.Error("Error on commit: " + err.Error())
				}
				return
			}
		}
	}()
}
