package productservice

import (
	"clean-code-structure/entity"
	"clean-code-structure/param/productparam"
	"clean-code-structure/repository/pgsql"
	"clean-code-structure/service/transactionservice"
	"context"
)

type Service struct {
	productRepo    ProductRepository
	transactionSvc transactionservice.Service
}

type ProductRepository interface {
	GetProductWithId(ctx context.Context, id int) (productparam.ProductRepo, error)
	StoreWitTransaction(ctx context.Context, transaction pgsql.Transaction, product entity.Product) (uint, error)
}

func New(
	productRepo ProductRepository,
	transactionSvc transactionservice.Service,
) Service {
	return Service{
		productRepo:    productRepo,
		transactionSvc: transactionSvc,
	}
}
