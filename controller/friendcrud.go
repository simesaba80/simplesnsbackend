package controller

import (
	"net/http"

	snsdb "snsback/db"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Sendfollow(c echo.Context) error {
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
	frienduser := snsdb.User{}
	if err := snsdb.DB.Model(&snsdb.User{}).Where("user_id = ?", obj.FriendId).First(&frienduser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Friend User Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	}
	friend := snsdb.Friend{}
	snsdb.DB.Table("friends").Where("user_id = ?", user.Id).Where("friend_id = ?", frienduser.Id).First(&friend)
	//フォローしてなかったらフォローする
	if friend.UserID == 0 && friend.FriendID == 0 {
		friend.UserID = user.Id
		friend.FriendID = frienduser.Id
		snsdb.DB.Create(&friend)
		return c.JSON(http.StatusOK, friend)
	}
	return c.JSON(http.StatusConflict, echo.Map{
		"message": "You already followed",
	})
}

func GetFriends(c echo.Context) error {
	userid := c.Param("userid")
	user := snsdb.User{}
	if err := snsdb.DB.Where("user_id = ?", userid).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Friend User Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	}
	friends := []snsdb.Friend{}
	snsdb.DB.Where("user_id = ?", user.Id).Find(&friends)
	return c.JSON(http.StatusOK, friends)
}
