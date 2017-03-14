package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"realworld/Model"
	"strconv"
)

func CheckUserLogin(c echo.Context) error {
	methodSource := "MethodSource : CheckUserLogin."
	jsonBody, errParse := parseJson(c)
	if !errParse {
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusBadRequest, "Failed To Parse Request")
	}
	exists, user := userLoginExists(jsonBody)
	response := new(Model.StandardResponse)
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
func FetchInterests(c echo.Context) error {
	methodSource := "MethodSource : fetchInterests."
	jsonBody, errParse := parseJson(c)
	var message = ""
	var statusCode = int64(200)
	response := new(Model.StandardResponse)
	if !errParse {
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusBadRequest, "Failed To Parse Request")
	}
	success, interests := getUserInterest(jsonBody["uid"])
	if success {
		message += "Fetched Interests Successfully"
		response.Data = interests
	} else {
		message += "Failed to fetch Interests"
		statusCode = 900
	}
	response.StatusCode = statusCode
	response.Message = message
	response.Success = success
	return c.JSON(http.StatusOK, response)
}

func FetchSimilarUsers(c echo.Context) error {
	methodSource := "MethodSource : fetchInterests."
	jsonBody, errParse := parseJson(c)
	var message = ""
	var statusCode = int64(200)
	response := new(Model.StandardResponse)
	if !errParse {
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusBadRequest, "Failed To Parse Request")
	}
	lat,_ := strconv.ParseFloat(jsonBody["lat"],64)
	lon,_ := strconv.ParseFloat(jsonBody["lat"],64)
	skip,_ := strconv.ParseInt(jsonBody["skip"],10,64)
	limit,_:=strconv.ParseInt(jsonBody["limit"],10,64)
	success,similarUsers := findSimilarUsers(
						jsonBody["uid"],
						lat,
						lon,
						skip,
						limit)
	if success{
		message+= "Fetched Similar Users Successfully"
		response.Data=similarUsers
	}else {
		message+= "Failed to fetch similar users"
		statusCode=900
	}
	response.StatusCode=statusCode
	response.Message=message
	response.Success=success
	return c.JSON(http.StatusOK,response)

}