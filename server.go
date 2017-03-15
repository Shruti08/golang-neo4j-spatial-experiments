package main

import (
	"net/http"
	"github.com/labstack/echo"
	"realworld/controllers"
	_ "gopkg.in/cq.v1"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

func getMeIn(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO FROM API")
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/letsGetIn", getMeIn)
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))


	r.POST("/createUser", controllers.CreateUser)
	r.POST("/checkUserLogin", controllers.CheckUserLogin)
	r.POST("/addUserInterests", controllers.CreateAddInterests)
	r.POST("/getUserInterests", controllers.FetchInterests)
	r.POST("/sendConnectionReq", controllers.SendConnectionRequest)
	r.POST("/acceptConnectionReq", controllers.AcceptConnectionRequest)
	r.POST("/blockUser", controllers.BlockUser)
	r.POST("/unblockUser", controllers.UnBlockUser)
	r.POST("/getSimilarUsers", controllers.FetchSimilarUsers)
	e.Logger.Fatal(e.Start(":8000"))
}