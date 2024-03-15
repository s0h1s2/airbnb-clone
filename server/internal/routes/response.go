package routes

type okApiResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}
type errorApiResponse struct {
	StatusCode int         `json:"statusCode"`
	Errors     interface{} `json:"errors"`
}
