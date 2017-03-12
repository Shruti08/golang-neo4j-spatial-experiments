package controllers
import ("time")
func relationshipProperty()map[string]string{
	var properties map[string]string
	properties = make(map[string]string)
	properties["createdOn"] = time.Now().String()
	return  properties
}
