package router

import (
	"net/http"
	"my_gorm/db"
	dbRepo "my_gorm/db/repository"
	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	ct := db.Connect()
	users , _ := dbRepo.GetListUsers(ct)
	return c.JSON(http.StatusOK, *users)
}

func Init(){
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", getUser)
	e.Logger.Fatal(e.Start(":1323"))
}

