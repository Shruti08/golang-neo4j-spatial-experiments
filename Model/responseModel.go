package Model

type StandardResponse struct {
	StatusCode int64 `json:"statusCode"`
	Success    bool `json:"success"`
	Message    string `json:"message"`
	Data       interface{} `json:"data"`
}
