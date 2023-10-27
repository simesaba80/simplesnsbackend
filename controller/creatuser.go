package controller

import (
	"net/http"

	snsdb "snsback/db"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	type Body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"Password"`
	}
	obj := Body{}
	user := snsdb.User{}
	if err := c.Bind(&obj); err != nil {
		//return 400
		return echo.NewHTTPError(http.StatusBadRequest, "BadRequest")
	}
	user.Name = obj.Name
	user.Email = obj.Email
	user.Password = obj.Password
	snsdb.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
