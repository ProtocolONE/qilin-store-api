package interfaces

import (
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
)

type AccountService interface {
	Authorize(userId string, authorize *dto.AuthorizeAccountDTO) (*model.User, error)
	Register(userId string, register *dto.RegisterAccountDTO) (*model.User, error)
}
