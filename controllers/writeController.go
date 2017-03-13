package controllers

import (
	"github.com/satori/go.uuid"
	"net/http"
	"realworld/Model"
	"github.com/mitchellh/mapstructure"
	"github.com/labstack/echo"
)
func CreateUser(c echo.Context) error {
	methodSource := "MethodSource : CreateUser."
	jsonBody, errParse := parseJson(c)
	var message string
	success := true
	if !errParse {
		return c.JSON(http.StatusBadRequest, "Failed To Parse Request")
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
		if createUserNode(jsonBody) {
			message += "User Created Successfully !"
			logMessage(message)
		} else {
			logMessage("NODE CREATION FAILED")
			return c.JSON(http.StatusInternalServerError, jsonBody)
		}
		logMessage(methodSource+"NEW ID " + u2.String())
	}
	response := new(Model.SingleUserResponse)
	response.StatusCode = statusCode
	response.Success = success
	response.Message = message
	if (success) {
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
	if (!parsed) {
		logMessage(methodSource + "Error Parsing User Interest JSON.")
	}

	for _, interest := range user.Interests {
		created := createInterestNode(interest)
		if (!created) {
			logMessage(methodSource + "Error Creating Interest Node.")
			statusCode = 201
			message += "\nError Creating Interest Node :" + interest
			success = false

		}
		added := createInterestRelationship(user.Uid, interest)
		if (!added) {
			logMessage(methodSource + "Error Adding Relationship.")
			statusCode = 201
			message += "\nError Adding To Relationship :" + interest
			success = false
		}
	}
	if message == "" {
		message = "Successfully added Interests !!!"
	}
	response := new(Model.SingleUserResponse)
	response.StatusCode = statusCode
	response.Success = success
	response.Message = message
	return c.JSON(http.StatusOK, response)
}
func AcceptConnection(c echo.Context) {


}
func sendRequest(c echo.Context){

}
func blockUser(c echo.Context){

}
func unBlockUser(c echo.Context){

}