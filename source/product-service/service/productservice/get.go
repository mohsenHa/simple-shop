package productservice

import (
	"clean-code-structure/param/productparam"
	"clean-code-structure/pkg/richerror"
)

func (s Service) Get(req productparam.GetRequest) (productparam.GetResponse, error) {
	const op = "productservice.Get"

	product, errGPWI := s.productRepo.GetProductWithId(req.Ctx, req.ProductId)
	if errGPWI != nil {
		return productparam.GetResponse{}, richerror.New(op).WithErr(errGPWI).
			WithMeta(map[string]interface{}{"product_id": req.ProductId})
	}

	return productparam.GetResponse{
		ProductId: product.Entity.Id,
		Name:      product.Entity.Name,
		Quantity:  product.Entity.Quantity,
	}, nil

}
