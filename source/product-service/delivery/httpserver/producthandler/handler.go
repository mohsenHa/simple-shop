package producthandler

import (
	"clean-code-structure/service/productservice"
	"clean-code-structure/validator/productvalidator"
)

type Handler struct {
	productValidator productvalidator.Validator
	productService   productservice.Service
}

func New(productService productservice.Service, productValidator productvalidator.Validator) Handler {
	return Handler{

		productService:   productService,
		productValidator: productValidator,
	}
}
