package controllers

import (
	"database/sql"
	"time"
)

func connectUsers(uid1 string, uid2 string) bool {
	methodSource := "MethodSource : connectUsers."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (a:User {uid:{0}}),(b:User{uid:{1}})
				 MERGE (a)-[r:CONNECTED]->(b)
				 SET r={2}
				 `)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(uid1, uid2, relationshipProperty())
	if errExec != nil {
		logMessage(methodSource + "Error executing query for Connection creation.Desc: " + errExec.Error())
		return false
	}
	return true
}

func createConnectionRequest(uid1 string, uid2 string) bool {
	methodSource := "MethodSource : createConnectionRequest."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (a:User {uid:{0}}),(b:User{uid:{1}})
				 MERGE (a)-[r:CONN_REQ ]->(b)
				 SET r={2}
				 `)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(uid1, uid2, relationshipProperty())
	if errExec != nil {
		logMessage(methodSource + "Error executing query for Connection creation.Desc: " + errExec.Error())
		return false
	}
	return true
}

func checkConnectionRequest(uid1 string, uid2 string) (bool, bool) {
	methodSource := "MethodSource : checkConnectionRequest."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false, false
	}
	defer db.Close()
	stmt, err := db.Prepare(`
				MATCH (n:User {uid: {0}})<-[r:CONN_REQ]-(m:User {uid: {1}})
				RETURN SIGN(COUNT(r))
				`)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false, false
	}
	defer stmt.Close()

	rows, errExec := stmt.Query(uid1, uid2)
	if errExec != nil {
		logMessage(methodSource + "Error executing query for Connection creation.Desc: " + errExec.Error())
		return false, false
	}
	var count = int64(-1)
	for rows.Next() {
		errScanner := rows.Scan(&count)
		if errScanner != nil {
			logMessage(methodSource + "Error Checking for connection request.Desc: " + errScanner.Error())
			return false, false
		}
	}
	if count == -1 {
		return false, false
	} else if count == 0 {
		return false, true
	} else {
		return true, true
	}
}

func blockUser(uid1 string, uid2 string) bool {
	methodSource := "MethodSource : blockUser."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (a:User {uid:{0}}),(b:User{uid:{1}})
				 MERGE (a)-[r:BLOCKED ]->(b)
				 SET r={2}
				 `)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(uid1, uid2, relationshipProperty())
	if errExec != nil {
		logMessage(methodSource + "Error executing query for Blocking Users.Desc: " + errExec.Error())
		return false
	}
	logMessage(uid1 + " blocked " + uid2 + " on " + time.Now().String())
	return true
}

func unblockUser(uid1 string, uid2 string) bool {
	methodSource := "MethodSource : blockUser."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (a:User {uid:{0}})-[r:BLOCKED]->(b:User{uid:{1}})
				 DELETE r
				 `)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(uid1, uid2, relationshipProperty())
	if errExec != nil {
		logMessage(methodSource + "Error executing query for Blocking Users.Desc: " + errExec.Error())
		return false
	}
	logMessage(uid1 + " unblocked " + uid2 + " on " + time.Now().String())
	return true
}
