package controllers

import (
	"log"
	"github.com/labstack/echo"
	"io/ioutil"
	"encoding/json"
)

func logMessage(logMsg string) {
	log.Printf(logMsg)
}

func parseJson(c echo.Context) (map[string]string, bool) {
	methodSource := " MethodSource : parseJson."
	s, errRead := ioutil.ReadAll(c.Request().Body)
	if errRead != nil {

		logMessage(methodSource + "Error while reading from request.Desc: " + errRead.Error())
		return map[string]string{}, false
	}
	jsonBody := map[string]string{}
	errParse := json.Unmarshal([]byte(s), &jsonBody)
	if errParse != nil {
		logMessage(methodSource + "Error while Parsing to Json. Desc: " + errParse.Error())
		return map[string]string{}, false
	}
	return jsonBody, true;
}