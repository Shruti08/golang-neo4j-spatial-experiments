package controllers

import (
	"database/sql"

	"golang-neo4j-spatial-experiments/Model"
)

func getBlockedUsers(uid string) (bool, []Model.BlockedUser) {
	methodSource := "MethodSource : getBlockedUsers."

	var blockedUsers []Model.BlockedUser
	db, err := sql.Open("neo4j-cypher", getConnectionUrl())
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false, blockedUsers
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (a:User{uid:{0}})-[:BLOCKED]->(b:User)
				RETURN b.uid,b.name,b.profilePicture`)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false, blockedUsers
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid)
	for rows.Next() {
		user := new(Model.BlockedUser)
		errScanner := rows.Scan(&user.Uid, &user.Name, &user.ProfilePicture)
		if errScanner != nil {
			logMessage(methodSource + "Error Checking for User.Desc: " + errScanner.Error())
			return false, blockedUsers
		}
		blockedUsers = append(blockedUsers, *user)
	}
	return true, blockedUsers
}

func getUserBuddies(uid string) (bool, []Model.ConnectedUser) {
	methodSource := "MethodSource : fetchInterests."
	var connectedUsers []Model.ConnectedUser
	db, err := sql.Open("neo4j-cypher", getConnectionUrl())
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false, connectedUsers
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (a:User{uid:{0}})-[:CONNECTED]-(b:User)
				RETURN b.uid,b.name,b.profilePicture`)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false, connectedUsers
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid)
	for rows.Next() {
		user := new(Model.ConnectedUser)
		errScanner := rows.Scan(&user.Uid, &user.Name, &user.ProfilePicture)
		if errScanner != nil {
			logMessage(methodSource + "Error Checking for User.Desc: " + errScanner.Error())
			return false, connectedUsers
		}
		successInterests, interests := getUserInterest(uid)
		if successInterests {
			user.Interests = interests.Interests
		} else {
			logMessage("Error fetching interests for user :" + user.Uid)
		}
		connectedUsers = append(connectedUsers, *user)
	}
	return true, connectedUsers
}
