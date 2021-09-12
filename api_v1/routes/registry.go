package routes

import (
	"github.com/stvkoch/copper/application/customer"
	"github.com/stvkoch/copper/application/subscription"

	"github.com/labstack/echo/v4"
)

type (
	handler interface {
		Routes(*echo.Group)
	}
)

func Routes(api *echo.Group) {
	handlers := []handler{
		customer.Handler{},
		subscription.Handler{},
	}

	for _, h := range handlers {
		h.Routes(api)
	}
}
