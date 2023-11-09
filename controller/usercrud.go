package controller

import (
	"net/http"

	snsdb "snsback/db"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	user := snsdb.User{}
	if err := snsdb.DB.Where("email = ?", obj.Email).First(&user).Error; err != nil {
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
