package controllers
import (
	"github.com/labstack/echo"
	"net/http"
	"realworld/Model"
)
func CheckUserLogin(c echo.Context) error {
	methodSource := "MethodSource : CheckUserLogin."
	jsonBody, errParse := parseJson(c)
	if !errParse {
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusBadRequest, "Failed To Parse Request")
	}
	exists, user := userLoginExists(jsonBody)
	response := new(Model.SingleUserResponse)
	if exists {
		response.StatusCode = 200
		response.Message = "User Already Exists - Logged In !"
		response.Success = true
		response.Data = *user
		return c.JSON(http.StatusOK, response)
	}
	response.StatusCode = 201
	response.Message = "New User"
	response.Success = true
	return c.JSON(http.StatusOK, response)
}
func FetchInterests(c echo.Context)error{
	methodSource := "MethodSource : fetchInterests."
	jsonBody, errParse := parseJson(c)
	var message =""
	var statusCode = int64(200)
	response := new(Model.SingleUserResponse)
	if !errParse {
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusBadRequest, "Failed To Parse Request")
	}
	success,interests := getUserInterest(jsonBody["uid"])
	if(success){
		message += "Fetched Interests Successfully"
		response.Data=interests
	}else{
		message += "Failed to fetch Interests"
		statusCode=900
	}
	response.StatusCode=statusCode
	response.Message=message
	response.Success=success
	return c.JSON(http.StatusOK,response)
}