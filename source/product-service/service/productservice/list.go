package productservice

import (
	"clean-code-structure/param/productparam"
	"clean-code-structure/pkg/richerror"
)

func (s Service) List(req productparam.ListRequest) (productparam.ListResponse, error) {
	const op = "productservice.List"

	products, hasMore, errGLP := s.productRepo.GetListProducts(req.Ctx, req.Page, req.PrePage)
	if errGLP != nil {
		return productparam.ListResponse{}, richerror.New(op).WithErr(errGLP)
	}

	return productparam.ListResponse{
		Products: products,
		HasMore:  hasMore,
		Page:     req.Page,
		PrePage:  req.PrePage,
	}, nil

}
