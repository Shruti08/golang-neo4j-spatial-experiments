package main

import (
	"database/sql"
	"net/http"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"encoding/json"
	_ "gopkg.in/cq.v1"
)

func logMessage(logMsg string){
	log.Printf(logMsg)
}

func parseJson(c echo.Context) (map[string]interface{},bool) {

	methodSource := " MethodSource : parseJson."

	s,errRead := ioutil.ReadAll(c.Request().Body)
	if errRead!=nil {
		logMessage(methodSource+"Error while reading from request.Desc: "+errRead.Error())
		return map[string]interface{}{},false
	}

	jsonBody := map[string]interface{}{}
	errParse := json.Unmarshal([]byte(s), &jsonBody)
	if errParse != nil {
		logMessage(methodSource+"Error while Parsing to Json. Desc: "+errParse.Error())
		return map[string]interface{}{},false
	}
	return  jsonBody,true;
}

func createNode(m map[string]interface{}) bool {
	db,err := sql.Open("neo4j-cypher","neo4j:srswios@123@http://localhost:7474")
	if err != nil {
		logMessage(err.Error())
		return false
	}
	defer db.Close()
	return  true
}

func createUser(c echo.Context) error{
	jsonBody,errParse := parseJson(c)
	if !errParse{
		return c.JSON(http.StatusBadRequest,"Failed To Parse Request")
	}
	u2 := uuid.NewV4()
	if createNode(jsonBody) {
		logMessage("NODE CREATED SUCCESSFULLY")
	} else {
		logMessage("NODE CREATION FAILED")
	}
	logMessage("NEW ID "+u2.String())
	return c.JSON(http.StatusOK,jsonBody)

}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Request().Body)
	})

	e.POST("/users",createUser)
	e.Logger.Fatal(e.Start(":8000"))
}