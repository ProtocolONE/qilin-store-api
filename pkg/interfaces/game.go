package interfaces

import "super_api/pkg/model"

type GameService interface {
	GetById(id string) (*model.Game, error)
	GetListGames(search string, offset int, limit int, order string) ([]*model.Game, error)
}