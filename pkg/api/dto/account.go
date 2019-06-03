package dto

import "time"

type (
	AuthorizeAccountDTO struct {
		Social *RegisterSocialDTO `json:"social"`
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

	UpdateUserDTO struct {
		Personal UpdatePersonalDTO `json:"personal"`
		Account  UpdateAccountDTO  `json:"account"`
	}

	UpdatePersonalDTO struct {
		FirstName string         `json:"first_name"`
		LastName  string         `json:"last_name"`
		BirthDate *time.Time     `json:"birth_date"`
		Address   UserAddressDTO `json:"address"`
	}

	UpdateAccountDTO struct {
		Nickname            string   `json:"nickname"`
		PrimaryLanguage     string   `json:"primary_language"`
		AdditionalLanguages []string `json:"additional_languages"`
	}

	LinkAccountDTO struct {

	}
)
