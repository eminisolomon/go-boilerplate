package dtos

type RegisterDto struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Lastname  string `json:"lastname" validate:"required,min=3"`
	Firstname string `json:"firstname" validate:"required,min=3"`
	Username  string `json:"username" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
}

type LoginDto struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type ChangePasswordDto struct {
	Password        string `json:"password" validate:"required"`
	NewPassword     string `json:"newpassword" validate:"required"`
	ConfirmPassword string `json:"confirmpassword" validate:"required"`
}
type ForgetPasswordDto struct {
	Email string `json:"email" validate:"required,email"`
}
