package handler

import "github.com/Toppira-Official/backend/internal/shared/entities"

type SignUpWithEmailPasswordInput struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=8,max=72" json:"password"`
}

func (in *SignUpWithEmailPasswordInput) MapUser() *entities.User {
	return &entities.User{
		Email:    in.Email,
		Password: &in.Password,
	}
}
