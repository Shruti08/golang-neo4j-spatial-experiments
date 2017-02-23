package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/satori/go.uuid"
	"database/sql"
)

func createNode(jsonBody map[string]string) bool {
	methodSource := " MethodSource : createNode."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()

	stmt, err := db.Prepare(`CREATE (user:User {0})
				 WITH count(*) AS dummy
				 MATCH(n:User) WHERE n.uid = {1} SET n.lat = toFloat(n.lat),n.lon=toFloat(n.lon)
				 WITH count(*) AS dummy
	                         MATCH (n:User) WHERE n.uid = {2} WITH n CALL spatial.addNode('geoLocation',n) YIELD node RETURN node;
	                         `)

	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Exec(jsonBody, jsonBody["uid"], jsonBody["uid"])

	if err != nil {
		logMessage(methodSource + "Error Adding Parameters.Desc: " + err.Error())
		return false
	}
	//defer rows.Close()
	rowsAffected, err := rows.RowsAffected()
	lastInsertId, err := rows.LastInsertId()
	logMessage("Rows Affected: " + string(rowsAffected) + ".Last Insert Id: " + string(lastInsertId))

	return true
}

func CreateUser(c echo.Context) error {
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