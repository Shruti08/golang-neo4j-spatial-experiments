package Model

type SingleUserResponse struct {
	StatusCode int64 `json:"statusCode"`
	Success    bool `json:"success"`
	Message    string `json:"message"`
	Data       User `json:"data"`
}
