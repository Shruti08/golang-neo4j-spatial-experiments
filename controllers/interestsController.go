package controllers

import ("realworld/Model"
	"database/sql"
)

func getUserInterest(uid string) (bool,Model.UserInterest){
	methodSource := "MethodSource : fetchInterests."
	userInterests := new(Model.UserInterest)
	userInterests.Uid=uid
	var interest string
	db, err := sql.Open("neo4j-cypher", "http://realworld:434Lw0RlD932803@localhost:7474")
	err = db.Ping()
	if err != nil {
		logMessage(methodSource + "Failed to Establish Connection. Desc: " + err.Error())
		return false,*userInterests
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
			return false,*userInterests
		}
		userInterests.Interests=append(userInterests.Interests,interest)
	}
	userInterests.Uid=uid
	return true,*userInterests
}
