package main

import (
	"os"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	// Echo instance
	e := echo.New()


	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Static("/", "web");


	if _, err := os.Stat(".keys/server.key"); errors.Is(err, os.ErrNotExist) {
		e.AutoTLSManager.Cache = autocert.DirCache(".cache")
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else {
		if err := e.StartTLS(":443", ".keys/server.crt", ".keys/server.key"); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}