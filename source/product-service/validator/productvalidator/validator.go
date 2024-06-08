package productvalidator

import "context"

type Validator struct {
	productRepo ProductRepository
}
type ProductRepository interface {
	IsProductExist(ctx context.Context, id uint) (bool, error)
}

func New(productRepo ProductRepository) Validator {
	return Validator{
		productRepo: productRepo,
	}
}
