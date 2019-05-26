package dto

import "time"

type (
	AuthorizeAccountDTO struct {
		Social    *RegisterSocialDTO `json:"social"`
	}

	RegisterAccountDTO struct {
		Email     string             `json:"email"`
		Birthdate *time.Time         `json:"birthdate"`
		Social    *RegisterSocialDTO `json:"social"`
	}

	RegisterSocialDTO struct {
		Provider string `json:"provider"`
		Token    string `json:"token"`
		Id       string `json:"id"`
	}
)
