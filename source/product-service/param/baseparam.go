package param

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseRequest struct {
	Ctx      context.Context
	Request  *http.Request
	Response *http.Response
}
type BaseResponse struct {
}

func NewBaseRequest(c echo.Context) BaseRequest {
	return BaseRequest{
		Ctx:      c.Request().Context(),
		Request:  c.Request(),
		Response: c.Request().Response,
	}
}
