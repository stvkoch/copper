package subscription

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Handler struct{}
)

func (handler Handler) Routes(e *echo.Group) {
	customerRouter := e.Group("/subscription")

	customerRouter.GET("/:id", handler.get)
	customerRouter.GET("", handler.all)
}

func (handler Handler) get(c echo.Context) error {
	return c.String(http.StatusOK, "get subs 1!\n")
}

func (handler Handler) all(c echo.Context) error {
	return c.String(http.StatusOK, "Get all subs 1!\n")
}
