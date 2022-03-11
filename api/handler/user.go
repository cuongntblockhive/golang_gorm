package handler

import (
	"my_gorm/db"
	dbRepo "my_gorm/db/repository"
	"my_gorm/entity"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	ct := db.Connect()
	users, _ := dbRepo.GetListUsers(ct)
	return c.JSON(http.StatusOK, *users)
}

func CreateUser(c echo.Context) error {
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

func Login(c echo.Context) error {
	// email := c.FormValue("email")
	// password := c.FormValue("password")

	ct := db.Connect()
	authUserData := entity.User{}
	c.Bind(&authUserData)
	err := dbRepo.SelectAuthUser(ct, &authUserData)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, entity.Error{Status: http.StatusUnauthorized, Message: "Fail"})
	}
	return c.JSON(http.StatusOK, authUserData)
}

