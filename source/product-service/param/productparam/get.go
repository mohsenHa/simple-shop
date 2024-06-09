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
	Product *ProductRepo `json:"product"`
}
