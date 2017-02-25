package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/satori/go.uuid"
	"database/sql"

	"github.com/labstack/gommon/log"
	"github.com/go-cq/cq/types"
)

func userLoginExists(json map[string]string) (bool, map[string]types.CypherValue) {
	methodSource := " MethodSource : userLoginExists."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false, nil
	}
	defer db.Close()

	stmt, err := db.Prepare(`MATCH (n:User)
			       WHERE (n.fbid = {0} OR n.gpid={1})
			       RETURN n
			       LIMIT 1`)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false, nil
	}

	defer stmt.Close()

	rows, err := stmt.Query(json["sid"], json["sid"])

	if err != nil {
		logMessage(methodSource + "Error executing query to check whether user exists.Desc: " + err.Error())
		return false, nil
	}
	defer rows.Close()

	var result map[string]types.CypherValue

	for rows.Next() {
		errScanner := rows.Scan(&result)
		if errScanner != nil {

			logMessage(methodSource + "Error Checking for User.Desc: " + errScanner.Error())
			return false, nil
		}
		logMessage("RESULT")
		log.Print(result)

	}

	if result == nil {
		return false, nil
	}
	return true, result

}
func CheckUserLogin(c echo.Context) error {
	methodSource := " MethodSource : CheckUserLogin."
	jsonBody, errParse := parseJson(c)
	if !errParse {
		logMessage(methodSource + "Error Parsing Request.")
		return c.JSON(http.StatusBadRequest, "Failed To Parse Request")
	}
	exists, user := userLoginExists(jsonBody)
	if exists {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusNotFound, "New User")
}

func userExists(json map[string]string) bool {

	return false;

}
func CreateUser(c echo.Context) error {
	jsonBody, errParse := parseJson(c)
	if !errParse {
		return c.JSON(http.StatusBadRequest, "Failed To Parse Request")
	}

	if userExists(jsonBody) {
		return c.JSON(http.StatusConflict, "User Already Exisits");
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
	return c.JSON(http.StatusCreated, jsonBody)

}