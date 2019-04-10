package services

import (
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"net/http"
	"super_api/pkg/common"
	"super_api/pkg/interfaces"
	"super_api/pkg/model"
)

type gameService struct {
	db interfaces.DatabaseProvider
}

const (
	CouldNotSaveGame = "Could not save game to db"
	CouldNotGetGame = "Could not retrieve game from db"
)

func NewGameService(db interfaces.DatabaseProvider) interfaces.GameService {
	return &gameService{db: db}
}

func (service *gameService) GetById(id string) (*model.Game, error) {
	db, err := service.db.GetDatabase()
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, errors.Wrap(err, common.GetDatabaseError))
	}

	result := model.Game{}
	err = db.C("games").Find(bson.M{"qilin_id": id}).One(&result)
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, errors.Wrap(err, CouldNotGetGame))
	}

	return &result, nil
}

func (gameService) GetListGames(search string, offset int, limit int, order string) ([]*model.Game, error) {
	panic("implement me")
}