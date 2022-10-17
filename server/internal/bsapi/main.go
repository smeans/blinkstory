package bsapi

import (
	"time"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Feed struct {
		TimeStamp time.Time
	}
)

func AddRoutes(group *echo.Group) {
	group.GET("/feed", getFeed)
}

func getFeed(c echo.Context) error {
	return c.JSON(http.StatusOK, &Feed{time.Now()})
}