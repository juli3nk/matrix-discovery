package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v10"
)

type ServerOptions struct {
	Port                  int      `default:"8080" validate:"port"`
	CORSAllowOrigins      []string `split_words:"true" default:"*"`
	CORSAllowMethods      []string `split_words:"true" default:"GET" validate:"cors_allow_methods"`
	MatrixHServer         string   `required:"true" split_words:"true" validate:"fqdn|hostname_port"`
	MatrixMHomeserver     string   `required:"true" split_words:"true" validate:"url"`
	MatrixMIdentityServer string   `split_words:"true" default:"https://vector.im" validate:"url"`
}

var opts *ServerOptions

func init() {
	opts = new(ServerOptions)

	if err := envconfig.Process("discovery", opts); err != nil {
		log.Fatal(err)
	}

	validate := validator.New()
	validate.RegisterValidation("cors_allow_methods", validateCorsAllowMethods)
	validate.RegisterValidation("port", validatePort)

	if err := validate.Struct(opts); err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: opts.CORSAllowOrigins,
		AllowMethods: opts.CORSAllowMethods,
	}))

	e.GET("/.well-known/matrix/server", getWellKnownMatrixServer)
	e.GET("/.well-known/matrix/client", getWellKnownMatrixClient)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", opts.Port)))
}

func getWellKnownMatrixServer(c echo.Context) error {
	wellKnown := map[string]string{"m.server": opts.MatrixHServer}

	return c.JSON(http.StatusOK, wellKnown)
}

func getWellKnownMatrixClient(c echo.Context) error {
	wellKnown := map[string]map[string]string{
		"m.homeserver":      map[string]string{"base_url": opts.MatrixMHomeserver},
		"m.identity_server": map[string]string{"base_url": opts.MatrixMIdentityServer},
	}

	return c.JSON(http.StatusOK, wellKnown)
}
