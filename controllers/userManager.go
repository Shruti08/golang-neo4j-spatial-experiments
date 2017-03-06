package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/satori/go.uuid"
	"database/sql"
	"github.com/labstack/gommon/log"
)

type singleUserResponse struct {
	StatusCode int64 `json:"statusCode"`
	Success    bool `json:"success"`
	Message    string `json:"Message"`
	Data       User `json:"data"`
}
type multiUserResponse struct {
	StatusCode int64 `json:"statusCode"`
	Success    bool `json:"success"`
	Message    string `json:"Message"`
	Data       []User `json:"data"`
}

type User struct {
	Name           string `json:"name"`
	Uid            string `json:"uid"`
	Fbid           string `json:"fbid"`
	Gpid           string `json:"gpid"`
	Email          string `json:"email"`
	Age            string `json:"age"`
	Dob            string `json:"dob"`
	Gender         string `json:"Gender"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	CreatedOn      string `json:"createdOn"`
	LastUpdateOn   string `json:"lastUpdateOn"`
	ProfilePicture string `json:"profilePicture"`
	DeviceToken    string `json:"deviceToken"`
	MobileNo       string `json:"mobileNo"`
}

func userLoginExists(json map[string]string) (bool, *User) {
	result := new(User)
	methodSource := " MethodSource : userLoginExists."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false, result
	}
	defer db.Close()

	stmt, err := db.Prepare(`MATCH (n:User)
			       WHERE (n.fbid = {0} OR n.gpid={1})
			       RETURN
			       n.name,
			       n.uid,
			       n.fbid,
			       n.gpid,
			       n.email,
			       n.age,
			       n.dob,
			       n.Gender,
			       n.lat,
			       n.lon,
			       n.createdOn,
			       n.lastUpdateOn,
			       n.profilePicture,
			       n.deviceToken,
			       n.mobileNo
			       LIMIT 1`)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false, result
	}

	defer stmt.Close()

	rows, err := stmt.Query(json["sid"], json["sid"])

	if err != nil {
		logMessage(methodSource + "Error executing query to check whether user exists.Desc: " + err.Error())
		return false, result
	}
	defer rows.Close()

	for rows.Next() {

		errScanner := rows.Scan(&result.Name,
			&result.Uid,
			&result.Fbid,
			&result.Gpid,
			&result.Email,
			&result.Age,
			&result.Dob,
			&result.Gender,
			&result.Lat,
			&result.Lon,
			&result.CreatedOn,
			&result.LastUpdateOn,
			&result.ProfilePicture,
			&result.DeviceToken,
			&result.MobileNo)
		if errScanner != nil {

			logMessage(methodSource + "Error Checking for User.Desc: " + errScanner.Error())
			return false, result
		}
		logMessage("RESULT")
		log.Print(result)

	}

	if result == nil || result.Uid == "" {
		return false, result
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
	response := new(singleUserResponse)
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
//TODO: COMPLETE THE FOLLOWING FOR USER CHECK DURING CREATION
func userExists(json map[string]string) bool {
	methodSource := " MethodSource : userExists."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (n:User)
			       WHERE (n.fbid = {0} OR n.gpid={1})
			       RETURN n
			       LIMIT 1`)

	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(json["fbid"], json["gpid"])

	for rows.Next() {
		return true;
	}

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