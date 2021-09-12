package main

import (
	"os"

	"github.com/go-playground/validator"
	"github.com/stvkoch/copper/lib"
	"github.com/stvkoch/copper/middlewares"
	"github.com/stvkoch/copper/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &lib.CustomValidator{Validator: validator.New()}

	apiV1 := e.Group("/api/v1")
	middlewares.Middlewares(e, apiV1)
	routes.Routes(apiV1)

	e.Logger.Fatal(e.Start(os.Getenv("API_SERVER_ADDR")))
}
