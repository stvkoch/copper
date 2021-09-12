package customer

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stvkoch/copper/application"
	"github.com/stvkoch/copper/domain"
)

type (
	Handler struct {
		application.Base
	}
	GetCustomerRequest struct {
		ID int `json:"id" param:"id" query:"id" validate:"numeric|required"`
	}
)

func (handler Handler) Routes(e *echo.Group) {
	customerRouter := e.Group("/customer")

	customerRouter.GET("/:id", handler.get)
	customerRouter.GET("", handler.all)
	customerRouter.PUT("", handler.update)
}

func (handler Handler) get(c echo.Context) error {
	params := &GetCustomerRequest{}
	return handler.BindValidate(c, params, func() (interface{}, error) {
		result, err := domain.GetCustomerByID(params.ID)
		return result, err
	})
}

func (handler Handler) all(c echo.Context) error {
	return c.String(http.StatusOK, "Get all customer 1!\n")
}

func (handler Handler) update(c echo.Context) error {
	return c.String(http.StatusOK, "update customer!\n")
}

func (handler Handler) create(c echo.Context) error {
	return c.String(http.StatusCreated, "created customer!\n")
}
