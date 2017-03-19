package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"realworld/Model"
	"strconv"
)

func CheckUserLogin(c echo.Context) error {
	methodSource := "MethodSource : CheckUserLogin."
	jsonBody, errParse := parseJson(c)
	response := new(Model.StandardResponse)
	if !errParse {
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
	}
	exists, user := userLoginExists(jsonBody)
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
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
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
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
	}
	lat, _ := strconv.ParseFloat(jsonBody["lat"], 64)
	lon, _ := strconv.ParseFloat(jsonBody["lon"], 64)
	skip, _ := strconv.ParseInt(jsonBody["skip"], 10, 64)
	limit, _ := strconv.ParseInt(jsonBody["limit"], 10, 64)
	success, similarUsers := findSimilarUsers(
		jsonBody["uid"],
		lat,
		lon,
		skip,
		limit)
	if success {
		message += "Fetched Similar Users Successfully"
		response.Data = similarUsers
	} else {
		message += "Failed to fetch similar users"
		statusCode = 900
	}
	response.StatusCode = statusCode
	response.Message = message
	response.Success = success
	return c.JSON(http.StatusOK, response)
}

func GetProfilePic(c echo.Context) error {
	id := c.Param("imageID")
	return c.File(id + ".png")
}

func FetchBuddies(c echo.Context) error {
	methodSource := "MethodSource : FetchBuddies."
	jsonBody, errParse := parseJson(c)
	var message = ""
	var statusCode = int64(200)
	response := new(Model.StandardResponse)
	if !errParse {
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
	}
	success, buddies := getUserBuddies(jsonBody["uid"])
	if success {
		message += "Fetched Buddies Successfully."
		response.Data = buddies
	} else {
		message += "Failed to fetch buddies."
		statusCode = 900
	}
	response.StatusCode = statusCode
	response.Message = message
	response.Success = success
	return c.JSON(http.StatusOK, response)
}

func FetchBlockedusers(c echo.Context) error {
	methodSource := "MethodSource : FetchBlockedusers."
	jsonBody, errParse := parseJson(c)
	var message = ""
	var statusCode = int64(200)
	response := new(Model.StandardResponse)
	if !errParse {
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
	}
	success, blockedUsers := getBlockedUsers(jsonBody["uid"])
	if success {
		message += "Fetched Blocked Users Successfully"
		response.Data = blockedUsers
	} else {
		message += "Failed to fetch blocked users"
		statusCode = 900
	}
	response.StatusCode = statusCode
	response.Message = message
	response.Success = success
	return c.JSON(http.StatusOK, response)
}

