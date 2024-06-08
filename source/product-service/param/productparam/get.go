package productparam

import (
	"clean-code-structure/param"
)

type GetRequest struct {
	param.BaseRequest
	ProductId int `param:"id"`
}

type GetResponse struct {
	param.BaseResponse
	ProductId int    `json:"product_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
}
