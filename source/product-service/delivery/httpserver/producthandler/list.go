package producthandler

import (
	"clean-code-structure/param"
	"clean-code-structure/param/productparam"
	"clean-code-structure/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) list(c echo.Context) error {
	req := productparam.ListRequest{
		BaseRequest: param.NewBaseRequest(c),
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if fieldErrors, err := h.productValidator.ValidateListRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	resp, err := h.productService.List(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)

}
