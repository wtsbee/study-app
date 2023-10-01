package logging

import (
	"log"

	"github.com/labstack/echo/v4"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエスト情報をロギング
		log.Printf("Request URI: %s, Method: %s", c.Request().RequestURI, c.Request().Method)

		return next(c)
	}
}
