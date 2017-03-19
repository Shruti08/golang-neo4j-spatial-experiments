package main

import (
	"github.com/labstack/echo"
	"realworld/controllers"
	_ "gopkg.in/cq.v1"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.GET("/profilePics/:imageID", controllers.GetProfilePic)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	r := e.Group("/api")
	r.POST("/createUser", controllers.CreateUser)
	r.POST("/checkUserLogin", controllers.CheckUserLogin)
	r.POST("/addUserInterests", controllers.CreateAddInterests)
	r.POST("/getUserInterests", controllers.FetchInterests)
	r.POST("/sendConnectionReq", controllers.SendConnectionRequest)
	r.POST("/acceptConnectionReq", controllers.AcceptConnectionRequest)
	r.POST("/blockUser", controllers.BlockUser)
	r.POST("/unblockUser", controllers.UnBlockUser)
	r.POST("/getSimilarUsers", controllers.FetchSimilarUsers)
	r.POST("/getUserBuddies", controllers.FetchBuddies)
	r.POST("/getBlockedUsers", controllers.FetchBlockedusers)
	e.Logger.Fatal(e.Start(":8000"))
}