package interfaces

import "github.com/ProtocolONE/qilin-store-api/pkg/model"

type GameService interface {
	GetById(id string) (*model.Game, error)
	GetListGames(search string, offset int, limit int, order string) ([]*model.Game, error)
}