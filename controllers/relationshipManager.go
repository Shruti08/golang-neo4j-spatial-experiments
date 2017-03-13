package controllers

import (
	"database/sql"
)

func createInterestRelationship(uid string, interest string) bool {
	methodSource := " MethodSource : createInterestRelationship."
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare(`MATCH (n:User {uid:{0}}),(i:Interest{name:{1}})
				 CREATE (n)-[r:LIKES {2}]->(i)
				 `)
	if err != nil {
		logMessage(methodSource + "Error Preparing Query.Desc: " + err.Error())
		return false
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(uid, interest, relationshipProperty())
	if errExec != nil {
		logMessage(methodSource + "Error executing query for Interest creation.Desc: " + errExec.Error())
		return false
	}
	return true
}
