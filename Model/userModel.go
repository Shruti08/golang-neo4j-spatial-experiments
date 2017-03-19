package Model

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

type SimilarUser struct {
	Name           string `json:"name"`
	Uid            string `json:"uid"`
	Gender         string `json:"Gender"`
	ProfilePicture string `json:"profilePicture"`
	Age            string `json:"age"`
	CreatedOn      string `json:"createdOn"`
	Interests      []string `json:"interests"`
	Distance       float64	`json:"distance"` 
	Connected      string `json:"connected"`
}

type BlockedUser struct {
	Name           string `json:"name"`
	Uid            string `json:"uid"`
	ProfilePicture string `json:"profilePicture"`
}

type ConnectedUser struct {
	Name           string `json:"name"`
	Uid            string `json:"uid"`
	ProfilePicture string `json:"profilePicture"`
	Interests      []string `json:"interests"`
}
