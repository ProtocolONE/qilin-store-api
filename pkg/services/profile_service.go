package services

import (
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/common"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"net/http"
)

type profileService struct {
	db interfaces.DatabaseProvider
}

func NewProfileService(db interfaces.DatabaseProvider) interfaces.ProfileService {
	return &profileService{db: db}
}

func (service *profileService) GetAccount(userId string) (*model.User, error) {
	db, err := service.db.GetDatabase()
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	account := &model.User{}
	err = db.C(accountsCollection).FindId(bson.ObjectIdHex(userId)).One(&account)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, common.NewServiceError(http.StatusNotFound, err)
		}
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	return account, nil
}

func (service *profileService) UpdateAccount(userId string, update dto.UpdateUserDTO) (*model.User, error) {
	db, err := service.db.GetDatabase()
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	account := &model.User{}
	id := bson.ObjectIdHex(userId)
	err = db.C(accountsCollection).FindId(id).One(account)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, common.NewServiceError(http.StatusNotFound, err)
		}
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	account.Account.Nickname = update.Account.Nickname
	account.Account.PrimaryLanguage = update.Account.PrimaryLanguage
	account.Account.AdditionalLanguages = update.Account.AdditionalLanguages
	account.Personal.Address.Region = update.Personal.Address.Region
	account.Personal.Address.PostalCode = update.Personal.Address.PostalCode
	account.Personal.Address.Line1 = update.Personal.Address.Line1
	account.Personal.Address.Line2 = update.Personal.Address.Line2
	account.Personal.Address.City = update.Personal.Address.City
	account.Personal.Address.Country = update.Personal.Address.Country
	account.Personal.BirthDate = update.Personal.BirthDate
	account.Personal.LastName = update.Personal.LastName
	account.Personal.FirstName = update.Personal.FirstName

	err = db.C(accountsCollection).UpdateId(id, account)
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	return account, nil
}


