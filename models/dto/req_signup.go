package dto

type ReqSignup struct {
	FullName        string `json:"fullName" validate:"required"`
	UserName        string `json:"userName" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6"`
	Gender          string `json:"gender" validate:"required,oneof=male female"`
}
