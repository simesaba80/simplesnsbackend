package main

import (
	"net/http"

	"snsback/controller"
	snsdb "snsback/db"

	"github.com/labstack/echo/v4"
)

func connect(c echo.Context) error {
	db, _ := snsdb.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB接続失敗しました")
	} else {
		return c.String(http.StatusOK, "DB接続しました")
	}
}

func main() {
	e := echo.New()
	e.GET("/", connect)
	//usercrud
	e.GET("/users/getusers", controller.GetUsers)
	e.GET("/users/getuser/:userid", controller.GetUser)
	e.POST("/users/creatuser", controller.CreateUser)
	e.PUT("/users/updateuser", controller.UpdateUser)
	e.POST("/users/login", controller.Login)
	//postcrud
	e.POST("/posts/createpost", controller.Createpost)
	e.GET("/posts/getposts", controller.GetPosts)
	e.GET("/posts/getpost/:id", controller.GetPost)
	e.PUT("/posts/updatepost/:id", controller.UpdatePost)
	e.DELETE("/posts/deletepost/:id", controller.DeletePost)
	//friendcrud
	e.POST("/friends/sendfollow", controller.Sendfollow)
	e.GET("/friends/getfriends/:userid", controller.GetFriends)
	e.DELETE("/friends/deletefollow", controller.DeleteFollow)
	e.Logger.Fatal(e.Start(":8080"))
}
