package controllers

import (
	"database/sql"
)

func createInterestRelationship(uid string, interest string) bool {
	methodSource := "MethodSource : createInterestRelationship."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (n:User {uid:{0}}),(i:Interest{name:{1}})
	                         MERGE (n)-[r:LIKES]->(i)
	                         SET r={2}
				 `)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(uid, interest,relationshipProperty())
	if errExec != nil {
		logMessage(methodSource + "Error executing query for Interest creation.Desc: " + errExec.Error())
		return false
	}
	return true
}

func createConnectionRelationship(uid1 string, uid2 string) bool {
	methodSource := "MethodSource : createConnectionRelationship."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (a:User {uid:{0}}),(b:User{uid:{1}})
				 CREATE (a)-[r:CONNECTED {2}]->(b)
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

func createRequestRealtionship(uid1 string, uid2 string) bool {
	methodSource := "MethodSource : createRequestRealtionship."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (a:User {uid:{0}}),(b:User{uid:{1}})
				 CREATE (a)-[r:CONN_REQ {2}]->(b)
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