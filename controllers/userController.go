package controllers

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	"realworld/Model"
)

func userLoginExists(json map[string]string) (bool, *Model.User) {
	result := new(Model.User)
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

func userExists(json map[string]string) (bool, string, int64, bool) {
	methodSource := " MethodSource : userExists."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false, "", 900, false
	}
	defer db.Close()
	var fbid, gpid, mobileNo, email string
	stmt, err := db.Prepare(`MATCH (n:User)
			       WHERE (n.fbid = {0} OR n.gpid={1} OR n.mobileNo={2} OR n.email={3})
			       RETURN n.fbid,n.gpid,n.mobileNo,n.email
			       LIMIT 1`)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false, "", 901, false
	}
	defer stmt.Close()

	rows, err := stmt.Query(json["fbid"], json["gpid"], json["mobileNo"], json["email"])

	for rows.Next() {
		errScanner := rows.Scan(&fbid, &gpid, &mobileNo, &email)
		if errScanner != nil {
			logMessage(methodSource + "Error Checking for User.Desc: " + errScanner.Error())
			return false, "", 902, false
		}
	}
	if (fbid != ""&&fbid == json["fbid"]) || (gpid != ""&&gpid == json["gpid"]) {
		return true, "social", 302, true
	} else if email != "" && email == json["email"] {
		return true, "email", 300, true
	} else if mobileNo != "" && mobileNo == json["mobileNo"] {
		return true, "mobileNo", 301, true
	}
	return false, "", 200, true;
}
