package productparam

import (
	"clean-code-structure/param"
)

type ListRequest struct {
	param.BaseRequest
	PrePage int `query:"pre_page"`
	Page    int `query:"page"`
}

type ListResponse struct {
	param.BaseResponse
	Products []ProductRepo `json:"products"`
	HasMore  bool          `json:"has_more"`
	PrePage  int           `json:"pre_page"`
	Page     int           `json:"page"`
}
