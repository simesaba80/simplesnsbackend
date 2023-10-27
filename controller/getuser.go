package controller

import (
	"net/http"

	snsdb "snsback/db"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := []snsdb.User{}
	snsdb.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	obj := Body{}
	user := snsdb.User{}
	if err := c.Bind(&obj); err != nil {
		return err
	}
	snsdb.DB.Where("Email = ?", obj.Email).Where("Password = ?", obj.Password).First(&user)

	return c.JSON(http.StatusOK, user)
}
