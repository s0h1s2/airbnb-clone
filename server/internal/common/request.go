package common

type OkApiResponse struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}
