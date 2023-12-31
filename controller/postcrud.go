package controller

import (
	"net/http"
	"time"

	"snsback/db"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Bodyをリクエストボディとして受け取り投稿を作成，IDは昇順
func CreatePost(c echo.Context) error {
	type Body struct {
		Content string `json:"content"`
		UserId  uint   `json:"userid"`
	}
	obj := Body{}
	post := db.Post{}
	user := db.User{}
	if err := c.Bind(&obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}
	db.DB.Where("id = ?", obj.UserId).First(&user)
	post.Content = obj.Content
	post.UserId = user.Id
	db.DB.Create(&post)
	return c.JSON(http.StatusCreated, post)
}

// 名前と内容を取得しスライスに格納しjsonで返す
func GetPosts(c echo.Context) error {
	//専用の返り値を宣言
	type Post struct {
		Content string `json:"content"`
		Name    string `json:"name"`
	}
	posts := []Post{}

	//postのUserIdをもとにテーブルを結合し投稿者の名前とpostの内容を取得するSQLをGORMで記述する
	//https://gorm.io/ja_JP/docs/query.html#Joins
	//https://gorm.io/ja_JP/docs/advanced_query.html
	if err := db.DB.Table("posts").Select("content, users.name").Joins("join users on posts.user_id = users.id").Order("posts.created_at DESC").Scan(&posts).Error; err != nil {
		//return 500
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, posts)
}

// パスパラメータで指定された投稿を取得
func GetPost(c echo.Context) error {
	//専用の返り値を宣言
	type Post struct {
		Content string `json:"content"`
		Name    string `json:"name"`
	}
	post := Post{}
	id := c.Param("id")

	//postのUserIdをもとにテーブルを結合し投稿者の名前とpostの内容を取得するSQLをGORMで記述する
	//https://gorm.io/ja_JP/docs/query.html#Joins
	//https://gorm.io/ja_JP/docs/advanced_query.html
	if err := db.DB.Table("posts").Select("content, users.name").Joins("join users on posts.user_id = users.id").Where("posts.id = ?", id).Order("posts.created_at DESC").Scan(&post).Error; err != nil {
		//return 500
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, post)
}

// パスパラメータで投稿を指定しリクエストボディをもとに更新
func UpdatePost(c echo.Context) error {
	type Body struct {
		Userid     uint   `json:"userid"`
		NewContent string `json:"newcontent"`
	}

	id := c.Param("id")
	obj := Body{}
	if err := c.Bind(&obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	post := db.Post{}
	db.DB.Table("posts").Where("id = ?", id).Find(&post)
	if post.Id == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error",
		})
	}
	if post.UserId != obj.Userid {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Unauthorized",
		})
	}
	post.Content = obj.NewContent
	post.UpdatedAt = time.Now()
	db.DB.Save(&post)
	return c.JSON(http.StatusOK, post)

}

// パスパラメータで投稿を指定しリクエストボディをもとに投稿を削除
func DeletePost(c echo.Context) error {
	type Body struct {
		UserId uint `json:"userid"`
	}
	id := c.Param("id")
	obj := Body{}
	if err := c.Bind(&obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}
	post := db.Post{}
	if err := db.DB.Where("id = ?", id).First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Challenge Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	}
	if post.UserId != obj.UserId {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Unauthorized",
		})
	}
	db.DB.Delete(&post)
	return c.JSON(http.StatusOK, post)
}
