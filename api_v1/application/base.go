package application

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Base struct{}

func (base Base) BindValidate(ctx echo.Context, params interface{}, handler func() (interface{}, error)) error {
	response := make(map[string]interface{})

	if err := ctx.Bind(params); err != nil {
		response["error"] = err
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	if err := ctx.Validate(params); err != nil {
		response["error"] = err
		return ctx.JSON(http.StatusBadRequest, response)
	}

	data, err := handler()

	response["data"] = data
	response["error"] = err

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response)
	}

	return ctx.JSON(http.StatusOK, response)
}
