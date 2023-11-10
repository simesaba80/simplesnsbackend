package controller

import (
	"net/http"
	"time"

	snsdb "snsback/db"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Bodyをリクエストボディとして受け取り投稿を作成，IDは昇順
func Createpost(c echo.Context) error {
	type Body struct {
		Content string `json:"content"`
		UserId  uint   `json:"userid"`
	}
	obj := Body{}
	post := snsdb.Post{}
	user := snsdb.User{}
	if err := c.Bind(&obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}
	snsdb.DB.Where("ID = ?", obj.UserId).First(&user)
	post.Content = obj.Content
	post.UserId = user.Id
	snsdb.DB.Create(&post)
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
	if err := snsdb.DB.Table("posts").Select("content, users.name").Joins("join users on posts.user_id = users.id").Order("posts.created_at DESC").Scan(&posts).Error; err != nil {
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
	if err := snsdb.DB.Table("posts").Select("content, users.name").Joins("join users on posts.user_id = users.id").Where("posts.id = ?", id).Order("posts.created_at DESC").Scan(&post).Error; err != nil {
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

	post := snsdb.Post{}
	snsdb.DB.Table("posts").Where("ID = ?", id).Find(&post)
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
	snsdb.DB.Save(&post)
	return c.JSON(http.StatusOK, post)

}

// パスパラメータで投稿を指定しリクエストボディをもとに投稿を削除
func DeletePost(c echo.Context) error {
	type Body struct {
		PostId uint `json:"postid"`
		UserId uint `json:"userid"`
	}
	obj := Body{}
	if err := c.Bind(&obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}
	post := snsdb.Post{}
	if err := snsdb.DB.Where("ID = ?", obj.PostId).First(&post).Error; err != nil {
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
	snsdb.DB.Delete(&post)
	return c.JSON(http.StatusOK, post)
}
