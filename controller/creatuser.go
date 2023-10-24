package controller

import (
	"net/http"

	snsdb "snsback/db"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := snsdb.User{}
	if err := c.Bind(&user); err != nil {
		//return 400
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	snsdb.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
