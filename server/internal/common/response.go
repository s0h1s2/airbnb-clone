package common

type ErrorApiResponse struct {
	Errors     interface{} `json:"errors"`
	StatusCode int         `json:"statusCode"`
}
