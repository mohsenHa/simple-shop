package productservice

import (
	"clean-code-structure/entity"
	"clean-code-structure/param/productparam"
	"clean-code-structure/pkg/richerror"
)

func (s Service) Store(req productparam.StoreRequest) (productparam.StoreResponse, error) {
	const op = "productservice.Store"

	transaction := s.transactionSvc.NewTransaction()

	id, errGPWI := s.productRepo.StoreWitTransaction(req.Ctx, transaction, entity.Product{
		Name: req.Name,
	})
	if errGPWI != nil {
		transaction.Rollback()
		return productparam.StoreResponse{}, richerror.New(op).WithErr(errGPWI)
	}
	transaction.Commit()

	return productparam.StoreResponse{
		ProductId: id,
		Name:      req.Name,
	}, nil
}
