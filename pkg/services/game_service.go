package services

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"net/http"
	"github.com/ProtocolONE/qilin-store-api/pkg/common"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
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
		if err == mgo.ErrNotFound {
			return nil, common.NewServiceError(http.StatusNotFound, errors.Wrap(err, CouldNotGetGame))
		}
		return nil, common.NewServiceError(http.StatusInternalServerError, errors.Wrap(err, CouldNotGetGame))
	}

	return &result, nil
}

func (service *gameService) GetListGames(search string, offset int, limit int, order string) ([]model.Game, error) {
	db, err := service.db.GetDatabase()
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, errors.Wrap(err, common.GetDatabaseError))
	}

	var games []model.Game
	err = db.C("games").Find(nil).All(&games)
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, errors.Wrap(err, "Could not get games list"))
	}

	return games, nil
}