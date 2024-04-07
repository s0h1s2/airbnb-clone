package common

type ErrorApiResponse struct {
	Errors     interface{} `json:"errors"`
	StatusCode int         `json:"statusCode"`
}

type OkApiResponse struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}
