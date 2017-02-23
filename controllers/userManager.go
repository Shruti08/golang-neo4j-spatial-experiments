package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/satori/go.uuid"
)


func CreateUser(c echo.Context) error {
	jsonBody, errParse := parseJson(c)
	if !errParse {
		return c.JSON(http.StatusBadRequest, "Failed To Parse Request")
	}
	u2 := uuid.NewV4()
	jsonBody["uid"] = u2.String()
	if createUserNode(jsonBody) {
		logMessage("NODE CREATED SUCCESSFULLY")
	} else {
		logMessage("NODE CREATION FAILED")
		return c.JSON(http.StatusInternalServerError, jsonBody)
	}
	logMessage("NEW ID " + u2.String())
	return c.JSON(http.StatusOK, jsonBody)

}