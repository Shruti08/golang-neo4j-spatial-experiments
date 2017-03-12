package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"realworld/Model"

)
func CreateAddInterests(c echo.Context) error {
	methodSource := " MethodSource : createAddInterests."
	user,parsed := parseJsonInterests(c)
	message := ""
	statusCode := int64(200)
	success := true
	if(!parsed){
		logMessage(methodSource+"Error Parsing User Interest JSON.")
	}

	for _, interest := range user.Interests {

		created := createInterestNode(interest)
		if (!created) {
			logMessage(methodSource + "Error Creating Interest Node.")
			statusCode = 201
			message +="\nError Creating Interest Node :"+interest
			success = false

		}

		added := createInterestRelationship(user.Uid, interest)
		if (!added) {
			logMessage(methodSource + "Error Adding Relationship.")
			statusCode = 201
			message +="\nError Adding To Relationship :"+interest
			success = false
		}
	}
	if message==""{
		message = "Successfully added Interests !!!"
	}
	response :=new(Model.SingleUserResponse)
	response.StatusCode = statusCode
	response.Success = success
	response.Message = message
	return c.JSON(http.StatusOK,response)
}