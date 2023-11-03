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

func GetPosts(c echo.Context) error {
	//専用の返り値を宣言
	type Post struct {
		Content string `json:"content"`
		Name    string `json:"name"`
	}
	posts := []Post{}
	//posts := []snsdb.Post{}

	//TODO: postのUserIdをもとにテーブルを結合し投稿者の名前とpostの内容を取得するSQLをGORMで記述する
	//https://gorm.io/ja_JP/docs/query.html#Joins
	//https://gorm.io/ja_JP/docs/advanced_query.html
	if err := snsdb.DB.Table("posts").Select("content, users.name").Joins("join users on posts.user_id = users.id").Order("users.created_at ASC").Scan(&posts).Error; err != nil {
		//return 500
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error: " + err.Error(),
		})
	}
	//snsdb.DB.Find(&posts)

	//postの配列をjsonとして返す
	return c.JSON(http.StatusOK, posts)
}
