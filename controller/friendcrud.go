package controller

import (
	"net/http"

	snsdb "snsback/db"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SendFriend(c echo.Context) error {
	type Body struct {
		UserId   string `json:"userid"`
		FriendId string `json:"friendid"`
	}
	obj := Body{}
	if err := c.Bind(&obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}
	user := snsdb.User{}
	if err := snsdb.DB.Model(&snsdb.User{}).Where("user_id = ?", obj.UserId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "User Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	}
	friend := snsdb.User{}
	if err := snsdb.DB.Model(&snsdb.User{}).Where("user_id = ?", obj.UserId).First(&friend).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "User Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	}
}
