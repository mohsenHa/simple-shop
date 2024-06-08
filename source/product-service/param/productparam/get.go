package productparam

import (
	"clean-code-structure/param"
)

type GetRequest struct {
	param.BaseRequest
	ProductId uint `param:"id"`
}

type GetResponse struct {
	param.BaseResponse
	ProductId uint   `json:"product_id"`
	Name      string `json:"name"`
}
