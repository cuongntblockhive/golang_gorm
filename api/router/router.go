package router

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"my_gorm/api/handler"

	"my_gorm/api/middleware"
)

func Init() {
	e := echo.New()

	// Middleware
	e.Use(middleware.CustomLogging())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", handler.GetUsers, middleware.Auth())
	e.POST("/user", handler.CreateUser)
	e.GET("/login", handler.Login)
	e.Logger.Fatal(e.Start(":1323"))
}
