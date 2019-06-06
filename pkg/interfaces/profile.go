package interfaces

import (
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
)

type ProfileService interface {
	GetAccount(userId string) (*model.User, error)
	UpdateAccount(userId string, update dto.UpdateUserDTO) (*model.User, error)
}
