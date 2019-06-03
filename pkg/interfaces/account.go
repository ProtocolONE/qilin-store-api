package interfaces

import (
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
)

type AccountService interface {
	Authorize(userId string, authorize dto.AuthorizeAccountDTO) (*model.User, error)
	Register(userId string, register dto.RegisterAccountDTO) (*model.User, error)
	GetAccount(userId string) (*model.User, error)
	UpdateAccount(userId string, update dto.UpdateUserDTO) (*model.User, error)
	RemoveMFA(userId string, providerId string) (*model.User, error)
	AddMFA(userId string, providerId string, providerName string) (*model.User, error)
}
