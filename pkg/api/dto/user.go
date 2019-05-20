package dto

import (
	"time"
)

type (
	UserDTO struct {
		ID       string                 `json:"id"`
		Personal PersonalInformationDTO `json:"personal"`
		Account  UserAccountDTO         `json:"account"`
		Security *UserSecurityDTO       `json:"security"`
	}

	PersonalInformationDTO struct {
		Email     string         `json:"email"`
		FirstName string         `json:"first_name"`
		LastName  string         `json:"last_name"`
		BirthDate *time.Time     `json:"birth_date"`
		Address   UserAddressDTO `json:"address"`
	}

	UserAddressDTO struct {
		Country    string `json:"country"`
		City       string `json:"city"`
		PostalCode string `json:"postal_code"`
		Region     string `json:"region"`
		Line1      string `json:"line_1"`
		Line2      string `json:"line_2"`
	}

	UserAccountDTO struct {
		Nickname            string                 `json:"nickname"`
		PrimaryLanguage     string                 `json:"primary_language"`
		AdditionalLanguages []string               `json:"additional_languages"`
		Socials             []UserSocialAccountDTO `json:"socials"`
	}

	UserSocialAccountDTO struct {
		Provider string `json:"provider"` // Facebook, twitter, Vk, etc.
		ID       string `json:"id"`
	}

	UserSecurityDTO struct {
		//TODO: дополнить при реализации Логина\Регистрации
	}
)
