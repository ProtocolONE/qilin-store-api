package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type (
	User struct {
		ID       bson.ObjectId       `json:"id" bson:"_id,omitempty"`
		Personal PersonalInformation `json:"personal"`
		Account  UserAccount         `json:"account"`
		Security *UserSecurity       `json:"security"`
	}

	PersonalInformation struct {
		Email     string      `json:"email"`
		FirstName string      `json:"first_name"`
		LastName  string      `json:"last_name"`
		BirthDate *time.Time  `json:"birth_date"`
		Address   UserAddress `json:"address"`
	}

	UserAddress struct {
		Country    string `json:"country"`
		City       string `json:"city"`
		PostalCode string `json:"postal_code"`
		Region     string `json:"region"`
		Line1      string `json:"line_1"`
		Line2      string `json:"line_2"`
	}

	UserAccount struct {
		Nickname            string              `json:"nickname"`
		PrimaryLanguage     string              `json:"primary_language"`
		AdditionalLanguages []string            `json:"additional_languages"`
		Socials             []UserSocialAccount `json:"socials"`
	}

	UserSocialAccount struct {
		//TODO: узнать какую именно информацию мы здесь держим. Токен, айди пользователя, ник?
		Provider string `json:"provider"` // Facebook, twitter, Vk, etc.
		ID       string `json:"id"`
	}

	UserSecurity struct {
		MFA []UserMFA `json:"mfa"`
	}

	UserMFA struct {
		ProviderId   string
		ProviderName string
	}
)
