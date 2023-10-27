package controller

import (
	"net/http"

	snsdb "snsback/db"

	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) error {
	type Body struct {
		Email      string `json:"Email"`
		Password   string `json:"password"`
		ReName     string `json:"rename"`
		ReEmail    string `json:"reemail"`
		RePassword string `json:"repassword"`
	}
	obj := Body{}
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "BadRequest")
	}
	var user snsdb.User
	snsdb.DB.Where("Email = ?", obj.Email).Where("Password = ?", obj.Password).First(&user)

	snsdb.DB.Model(&user).Updates(snsdb.User{Name: obj.ReName, Email: obj.ReEmail, Password: obj.RePassword})
	return c.JSON(http.StatusOK, user)
}
