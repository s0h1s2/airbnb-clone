package users

type createUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type userInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type loginUserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type userInfoResponse struct {
	User userInfo `json:"user"`
}
