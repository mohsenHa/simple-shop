package productparam

import (
	"clean-code-structure/param"
)

type StoreRequest struct {
	param.BaseRequest
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type StoreResponse struct {
	param.BaseResponse
	ProductId uint   `json:"product_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
}
