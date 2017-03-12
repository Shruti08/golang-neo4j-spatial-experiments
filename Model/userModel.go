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
