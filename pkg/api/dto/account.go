package dto

import "time"

type (
	AuthorizeAccountDTO struct {
		Social *RegisterSocialDTO `json:"social"`
	}

	RegisterAccountDTO struct {
		Email     string             `json:"email" validate:"required,email,max=254"`
		Nickname  string             `json:"nickname" validate:"max=50"`
		Birthdate *time.Time         `json:"birthdate"`
		Social    *RegisterSocialDTO `json:"social"`
	}

	RegisterSocialDTO struct {
		Provider string `json:"provider" validate:"max=254"`
		Token    string `json:"token" validate:"max=254"`
		Id       string `json:"id" validate:"max=254"`
	}

	UpdateUserDTO struct {
		Personal UpdatePersonalDTO `json:"personal"`
		Account  UpdateAccountDTO  `json:"account"`
	}

	UpdatePersonalDTO struct {
		FirstName string         `json:"first_name" validate:"max=50"`
		LastName  string         `json:"last_name" validate:"max=50"`
		BirthDate *time.Time     `json:"birth_date"`
		Address   UserAddressDTO `json:"address" validate:"max=254"`
	}

	UpdateAccountDTO struct {
		Nickname            string   `json:"nickname" validate:"max=50"`
		PrimaryLanguage     string   `json:"primary_language"`
		AdditionalLanguages []string `json:"additional_languages"`
	}

	UpdateMultiFactorAuthDTO struct {
		ProviderName string
		ProviderId   string
	}

	LinkAccountDTO struct {
	}
)
