package main

import (
	"net/http"
	"snsback/controller"
	"snsback/db"
	"snsback/utils/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.LoadEnv()
	db.Init()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//usercrud
	e.GET("/users/getusers", controller.GetUsers)
	e.GET("/users/getuser/:userid", controller.GetUser)
	e.POST("/users/createuser", controller.CreateUser)
	e.PUT("/users/updateuser", controller.UpdateUser)
	e.POST("/users/login", controller.Login)
	//postcrud
	e.POST("/posts/createpost", controller.CreatePost)
	e.GET("/posts/getposts", controller.GetPosts)
	e.GET("/posts/getpost/:id", controller.GetPost)
	e.PUT("/posts/updatepost/:id", controller.UpdatePost)
	e.DELETE("/posts/deletepost/:id", controller.DeletePost)
	//friendcrud
	e.POST("/friends/sendfollow", controller.SendFollow)
	e.GET("/friends/getfriends/:userid", controller.GetFollowList)
	e.DELETE("/friends/deletefollow", controller.DeleteFollow)
	e.Logger.Fatal(e.Start(":8080"))
}
