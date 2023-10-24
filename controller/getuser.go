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
