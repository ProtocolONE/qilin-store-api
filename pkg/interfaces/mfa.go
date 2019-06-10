package interfaces

import "github.com/ProtocolONE/qilin-store-api/pkg/api/dto"

type MfaService interface {
	Add(userId string, providerId string) error
	List(userId string) ([]dto.MfaProviderDTO, error)
	Remove(userId string, providerId string) error
}
