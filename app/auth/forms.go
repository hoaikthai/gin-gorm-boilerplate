package auth

type SignInForm struct {
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignUpForm struct {
	Email    string `json:"email" binding:"required,email"`
	UserName string `json:"username" binding:"required,max=20"`
	Password string `json:"password" binding:"required,max=20,min=8"`
}
