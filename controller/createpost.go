package controller

import (
	"net/http"

	snsdb "snsback/db"

	"github.com/labstack/echo/v4"
)

func Createpost(c echo.Context) error {
	type Body struct {
		Content string `json:"content"`
		UserId  uint   `json:"userid"`
	}
	obj := Body{}
	post := snsdb.Post{}
	user := snsdb.User{}
	if err := c.Bind(&obj); err != nil {
		//return 400
		return echo.NewHTTPError(http.StatusBadRequest, "BadRequest")
	}
	snsdb.DB.Where("ID = ?", obj.UserId).First(&user)
	post.Content = obj.Content
	post.UserId = user.ID
	snsdb.DB.Create(&post)
	return c.JSON(http.StatusCreated, post)
}
