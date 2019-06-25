package dto

type MfaProviderDTO struct {
	// ID is the id of provider.
	ID string `json:"id"`

	// AppID is the id of the application.
	AppID string `json:"app_id"`

	// Name is a human-readable name of provider.
	Name string `json:"name"`
}

type GetListMfaProviderDTO struct {
	ClientId string `json:"client_id" form:"client_id"`
}