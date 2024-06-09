package producthandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(messageGroup *echo.Group) {

	messageGroup.GET("", h.list)
	messageGroup.GET("/:id", h.get)

}
