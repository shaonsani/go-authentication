package request

import "mime/multipart"

type RegisterRequest struct {
	Name     string                `validate:"required,min=2,max=100" json:"name"`
	Email    string                `validate:"required,email" json:"email"`
	Gender   string                `validate:"required,oneof=male female other" json:"gender"`
	Mobile   string                `validate:"required,len=10" json:"mobile"`
	Address  string                `validate:"required,max=255" json:"address"`
	Photo    *multipart.FileHeader `validate:"omitempty" json:"photo"`
	Password string                `validate:"required,min=8" json:"password"`
}

type LoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}
