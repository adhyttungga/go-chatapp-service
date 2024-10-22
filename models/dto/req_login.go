package dto

type ReqLogin struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}
