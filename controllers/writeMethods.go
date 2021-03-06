package controllers

import (
	"net/http"
	"time"

	"golang-neo4j-spatial-experiments/Model"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
	"github.com/satori/go.uuid"
)

func CreateUser(c echo.Context) error {
	methodSource := "MethodSource : CreateUser."
	jsonBody, errParse := parseJson(c)
	var message string
	success := true
	response := new(Model.StandardResponse)
	if !errParse {
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
	}
	exists, field, statusCode, methodSuccess := userExists(jsonBody)
	if !methodSuccess {
		success = false
		statusCode = statusCode
		message = "Something Went Wrong."
	} else if exists {
		success = false
		statusCode = statusCode
		message = "User Already Exists. Duplicate " + field
	} else {
		u2 := uuid.NewV4()
		jsonBody["uid"] = u2.String()
		jsonBody["createdOn"], jsonBody["lastUpdateOn"] = time.Now().String(), time.Now().String()
		picSave := saveImage(u2.String(), jsonBody["pic"])
		if !picSave {
			logMessage("Error saving profile picture for " + u2.String())

		}
		jsonBody["pic"] = "/profilePics/" + u2.String() + ".png"
		if createUserNode(jsonBody) {
			message += "User Created Successfully !"
			logMessage(message)
		} else {
			logMessage("NODE CREATION FAILED")
			statusCode = 900
			success = false
			message = "Node Creation Failed.Invalid or Null values passed."
		}
		logMessage(methodSource + "NEW ID " + u2.String())
	}
	response.StatusCode = statusCode
	response.Success = success
	response.Message = message
	if success {
		user := new(Model.User)
		mapstructure.Decode(jsonBody, user)
		response.Data = *user
	}
	return c.JSON(http.StatusOK, response)
}

func CreateAddInterests(c echo.Context) error {
	methodSource := " MethodSource : createAddInterests."
	user, parsed := parseJsonInterests(c)
	message := ""
	statusCode := int64(200)
	success := true
	response := new(Model.StandardResponse)
	if !parsed {
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
	}
	if !emptyUserInterests(user.Uid) {
		response.StatusCode = 900
		response.Message = "Failed to Clear all user interests."
		response.Success = false
		logMessage(methodSource + "Failed to Clear all user interests.")
		return c.JSON(http.StatusOK, response)
	}
	for _, interest := range user.Interests {
		created := createInterestNode(interest)
		if !created {
			logMessage(methodSource + "Error Creating Interest Node.")
			statusCode = 201
			message += "\nError Creating Interest Node :" + interest
			success = false

		}
		added := addUserInterests(user.Uid, interest)
		if !added {
			logMessage(methodSource + "Error Adding Relationship.")
			statusCode = 201
			message += "\nError Adding To Relationship :" + interest
			success = false
		}
	}
	if message == "" {
		message = "Successfully added Interests !!!"
	}
	response.StatusCode = statusCode
	response.Success = success
	response.Message = message
	return c.JSON(http.StatusOK, response)
}

func AcceptConnectionRequest(c echo.Context) error {
	methodSource := "MethodSource : AcceptConnRequest."
	jsonBody, errParse := parseJson(c)
	var message = ""
	var statusCode = int64(200)
	var success = true
	response := new(Model.StandardResponse)
	if !errParse {
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
	}
	connExists, methodSuccess := checkConnectionRequest(jsonBody["uid1"], jsonBody["uid2"])
	if !methodSuccess {
		message += "Error checking for connection request."
		statusCode = int64(900)
		success = false

	} else if connExists {
		if !connectUsers(jsonBody["uid1"], jsonBody["uid2"]) {
			message += "Error creating connection link."
			statusCode = int64(900)
			success = false
		} else {
			message += "Successfully connected users."
		}
	} else {
		message += "Connection Request not found"
		statusCode = int64(900)
		success = false
	}
	response.Message = message
	response.Success = success
	response.StatusCode = statusCode
	return c.JSON(http.StatusOK, response)
}

func SendConnectionRequest(c echo.Context) error {
	methodSource := "MethodSource : SendConnRequest."
	jsonBody, errParse := parseJson(c)
	var message = ""
	var statusCode = int64(200)
	var success = true
	response := new(Model.StandardResponse)
	if !errParse {
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
	}
	requestSent := createConnectionRequest(jsonBody["uid1"], jsonBody["uid2"])
	if !requestSent {
		success = false
		message += "Unable to send Request."
		statusCode = 900
	}
	response.StatusCode = statusCode
	response.Message = message
	response.Success = success
	return c.JSON(http.StatusOK, response)
}

func BlockUser(c echo.Context) error {
	methodSource := "MethodSource : BlockUser."
	jsonBody, errParse := parseJson(c)
	var message = ""
	var statusCode = int64(200)
	var success = true
	response := new(Model.StandardResponse)
	if !errParse {
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
	}
	requestSent := blockUser(jsonBody["uid1"], jsonBody["uid2"])
	if !requestSent {
		success = false
		message += "Unable to block user.Internal server Error."
		statusCode = 900
	}
	response.StatusCode = statusCode
	response.Message = message
	response.Success = success
	return c.JSON(http.StatusOK, response)
}

func UnBlockUser(c echo.Context) error {
	methodSource := "MethodSource : UnBlockUser."
	jsonBody, errParse := parseJson(c)
	var message = ""
	var statusCode = int64(200)
	var success = true
	response := new(Model.StandardResponse)
	if !errParse {
		response.StatusCode = 900
		response.Message = "Failed to parse request. Invalid JSON"
		response.Success = false
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusOK, response)
	}
	requestSent := unblockUser(jsonBody["uid1"], jsonBody["uid2"])
	if !requestSent {
		success = false
		message += "Unable to unblock User.Internal server error."
		statusCode = 900
	}
	response.StatusCode = statusCode
	response.Message = message
	response.Success = success
	return c.JSON(http.StatusOK, response)
}
