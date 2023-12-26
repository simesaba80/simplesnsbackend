package controller

import (
	"net/http"
	"time"

	"snsback/db"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Bodyをもとにユーザーを作成
func CreateUser(c echo.Context) error {
	type Body struct {
		Name     string `json:"name"`
		UserID   string `json:"userid"`
		Email    string `json:"email"`
		Password string `json:"Password"`
	}
	obj := Body{}
	user := db.User{}
	if err := c.Bind(&obj); err != nil {
		//return 400
		return echo.NewHTTPError(http.StatusBadRequest, "BadRequest")
	}
	user.Name = obj.Name
	user.UserID = obj.UserID
	user.Email = obj.Email
	user.Password = obj.Password
	db.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

// 全ユーザーを取得
func GetUsers(c echo.Context) error {
	users := []db.User{}
	db.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// パスパラメータのユーザーIDをもとにユーザーを指定しBody型の構造体に入れて返す
func GetUser(c echo.Context) error {
	type Body struct {
		Name      string    `json:"name"`
		UserID    string    `json:"userid"`
		CreatedAt time.Time `json:"createdat"`
	}
	user := Body{}

	user.UserID = c.Param("userid")
	db.DB.Model(&db.User{}).Where("user_id = ?", user.UserID).First(&user)

	return c.JSON(http.StatusOK, user)
}

// リクエストボディのメアドとパスワードでユーザーを指定し、名前メアドパスワードを更新
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
	var user db.User
	db.DB.Where("email = ?", obj.Email).Where("password = ?", obj.Password).First(&user)

	db.DB.Model(&user).Updates(db.User{Name: obj.ReName, Email: obj.ReEmail, Password: obj.RePassword})
	return c.JSON(http.StatusOK, user)
}

// メアドとパスワードでユーザーを指定しIDを返す
func Login(c echo.Context) error {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	obj := Body{}
	if err := c.Bind(&obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}
	user := db.User{}
	if err := db.DB.Where("email = ?", obj.Email).First(&user).Error; err != nil {
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
	return c.JSON(http.StatusOK, user.Id)
}
