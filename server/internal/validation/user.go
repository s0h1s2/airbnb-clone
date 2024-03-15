package validation

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email" `
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,string,min=5,max=72"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email" `
	Password string `json:"password" binding:"required"`
}
