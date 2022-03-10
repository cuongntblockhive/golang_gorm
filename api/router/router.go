package router

import (
	// "fmt"
	// "fmt"
	"my_gorm/db"
	dbRepo "my_gorm/db/repository"
	"my_gorm/entity"

	"net/http"

	"github.com/labstack/echo/v4"
)

func getUsers(c echo.Context) error {
	ct := db.Connect()
	users, _ := dbRepo.GetListUsers(ct)
	return c.JSON(http.StatusOK, *users)
}

func createUser(c echo.Context) error {
	// name := c.FormValue("name")
	// email := c.FormValue("email")
	// age := c.FormValue("age")
	// newUser := entity.User{Name: name , Email: email, Age: age}

	newUser := entity.User{}
	err := c.Bind(&newUser)
	
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, entity.Error{Status: http.StatusNotAcceptable, Message: err.Error()})
	}

	ct := db.Connect()
	err = dbRepo.CreateUser(ct, &newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Error{Status: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, newUser)
}

func Init() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", getUsers)
	e.POST("/user", createUser)
	e.Logger.Fatal(e.Start(":1323"))
}
