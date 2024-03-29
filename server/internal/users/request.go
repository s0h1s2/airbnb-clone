package users

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email" `
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=5,max=72"`
}
type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email" `
	Password string `json:"password" binding:"required,min=5,max=72"`
}
