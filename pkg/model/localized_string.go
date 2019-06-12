package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
)

type LocalizedStringArray struct {
	// english
	EN []string `json:"en"`

	// russian
	RU []string `json:"ru,omitempty"`

	// other languages
	FR []string `json:"fr,omitempty"`
	ES []string `json:"es,omitempty"`
	DE []string `json:"de,omitempty"`
	IT []string `json:"it,omitempty"`
	PT []string `json:"pt,omitempty"`
}

func (p LocalizedStringArray) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return string(j), err
}

func (p *LocalizedStringArray) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}
	if err := json.Unmarshal(source, &p); err != nil {
		return err
	}
	return nil
}

func (p LocalizedStringArray) GetValueOrDefault(lng string) []string {
	var value []string
	switch strings.ToUpper(lng) {
	case "EN":
		value = p.EN
	case "RU":
		value = p.RU
	case "FR":
		value = p.FR
	case "ES":
		value = p.ES
	case "DE":
		value = p.DE
	case "IT":
		value = p.IT
	case "PT":
		value = p.PT
	}

	if len(value) == 0 {
		value = p.EN
	}
	return value
}

// LocalizedString is helper object to hold localized string properties.
type LocalizedString struct {
	// english
	EN string `json:"en"`

	// russian
	RU string `json:"ru,omitempty"`

	// other languages
	FR string `json:"fr,omitempty"`
	ES string `json:"es,omitempty"`
	DE string `json:"de,omitempty"`
	IT string `json:"it,omitempty"`
	PT string `json:"pt,omitempty"`
}

func (p LocalizedString) GetValueOrDefault(lng string) string {
	value := ""
	switch strings.ToUpper(lng) {
	case "EN":
		value = p.EN
	case "RU":
		value = p.RU
	case "FR":
		value = p.FR
	case "ES":
		value = p.ES
	case "DE":
		value = p.DE
	case "IT":
		value = p.IT
	case "PT":
		value = p.PT
	}

	if value == "" {
		value = p.EN
	}
	return value
}

func (p LocalizedString) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return string(j), err
}

func (p *LocalizedString) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}
	if err := json.Unmarshal(source, &p); err != nil {
		return err
	}
	return nil
}
