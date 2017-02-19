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

func createNode(jsonBody map[string]string) bool {
	methodSource := " MethodSource : createNode."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()

	stmt,err :=db.Prepare(`CREATE (user:User {0})`)


	if err!=nil{
		logMessage(methodSource+"Error Preparing Query.Desc: "+err.Error())
		return false
	}
	defer stmt.Close()


          rows,err := stmt.Exec(jsonBody)



	if err!=nil {
		logMessage(methodSource+"Error Adding Parameters.Desc: "+err.Error())
		return false
	}
	//defer rows.Close()
	rowsAffected,err:=rows.RowsAffected()
	lastInsertId,err:=rows.LastInsertId()
	logMessage("Rows Affected: "+string(rowsAffected)+".Last Insert Id: "+string(lastInsertId))



	return true
}

func createUser(c echo.Context) error {
	jsonBody, errParse := parseJson(c)
	if !errParse {
		return c.JSON(http.StatusBadRequest, "Failed To Parse Request")
	}
	u2 := uuid.NewV4()
	jsonBody["uid"] = u2.String()
	if createNode(jsonBody) {
		logMessage("NODE CREATED SUCCESSFULLY")
	} else {
		logMessage("NODE CREATION FAILED")
		return c.JSON(http.StatusInternalServerError, jsonBody)
	}
	logMessage("NEW ID " + u2.String())
	return c.JSON(http.StatusOK, jsonBody)

}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Request().Body)
	})

	e.POST("/users", createUser)
	e.Logger.Fatal(e.Start(":8000"))
}