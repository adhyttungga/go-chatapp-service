package dto

type ResUser struct {
	ID         string `json:"_id"`
	FullName   string `json:"fullName"`
	UserName   string `json:"userName"`
	ProfilePic string `json:"profilePic"`
}
