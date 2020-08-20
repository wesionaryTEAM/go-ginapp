package dto

//LoginCredentials ..
type LoginCredentials struct {
	Username string `form:"email"`
	Password string `form:"password"`
}
