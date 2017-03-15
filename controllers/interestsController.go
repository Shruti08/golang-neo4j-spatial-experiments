package controllers

import (
	"realworld/Model"
	"database/sql"
	"github.com/labstack/gommon/log"
)

func getUserInterest(uid string) (bool, Model.UserInterest) {
	methodSource := "MethodSource : fetchInterests."
	userInterests := new(Model.UserInterest)
	userInterests.Uid = uid
	var interest string
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false, *userInterests
	}
	defer db.Close()

	stmt, err := db.Prepare(`MATCH (n:User{uid:{0}})-[:LIKES]->(i:Interest)
				RETURN i.name`)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false, *userInterests
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid)

	for rows.Next() {
		errScanner := rows.Scan(&interest)
		if errScanner != nil {
			logMessage(methodSource + "Error Checking for User.Desc: " + errScanner.Error())
			return false, *userInterests
		}
		userInterests.Interests = append(userInterests.Interests, interest)
	}
	userInterests.Uid = uid
	return true, *userInterests
}

func findSimilarUsers(uid string, lat float64, lon float64, skip int64, limit int64) (bool, []Model.SimilarUser) {
	methodSource := "MethodSource : findSimilarUsers."
	var similarUsers []Model.SimilarUser
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false, similarUsers
	}
	defer db.Close()
	log.Print(uid, "#", lat, "#", lon, "#", skip, "#", limit)
	stmt, err := db.Prepare(`
				CALL spatial.withinDistance("geoLocation",{lat:{0},lon:{1}},7)
				YIELD node
				MATCH (node)-[:LIKES]->(x:Interest)<-[:LIKES]-(a:User{uid:{2}})
				OPTIONAL MATCH(node)-[r:CONNECTED]-(a:User{uid:{3}})
				RETURN  node.uid,
				 	node.name,
				 	node.Gender,
				 	node.profilePicture,
				 	node.age,
				 	node.lastUpdateOn AS dateTime,
				 	(CASE WHEN COUNT(r)>0 THEN 1 ELSE 0 END) AS CONNECTED
				ORDER BY dateTime  DESC
				SKIP {4}
				LIMIT {5}
				`)

	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false, similarUsers
	}
	rows, errExec := stmt.Query(lat, lon, uid, uid, skip, limit)

	if errExec != nil {
		logMessage(methodSource + "Error executing query for findig similar users.Desc: " + errExec.Error())
		return false, similarUsers
	}
	for rows.Next() {
		user := new(Model.SimilarUser)
		errScanner := rows.Scan(&user.Uid,
			&user.Name,
			&user.Gender,
			&user.ProfilePicture,
			&user.Age,
			&user.CreatedOn,
			&user.Connected)
		if errScanner != nil {
			logMessage(methodSource + "Error Checking for Similar Users.Desc: " + errScanner.Error())
			return false, similarUsers
		}
		log.Print(user)
		similarUsers = append(similarUsers, *user)

	}
	defer stmt.Close()
	return true, similarUsers

}