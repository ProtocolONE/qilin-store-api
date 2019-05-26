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

type accountService struct {
	db interfaces.DatabaseProvider
}

func NewAccountService(db interfaces.DatabaseProvider) interfaces.AccountService {
	return &accountService{db}
}

func (service *accountService) Authorize(userId string, authorize *dto.AuthorizeAccountDTO) (*model.User, error) {
	db, err := service.db.GetDatabase()
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	account := &model.User{}
	err = db.C("accounts").FindId(bson.ObjectIdHex(userId)).One(&account)
	if err != nil {
		if err == mgo.ErrNotFound{
			return nil, common.NewServiceError(http.StatusNotFound, err)
		}
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	return account, nil
}

func (service *accountService) Register(userId string, register *dto.RegisterAccountDTO) (*model.User, error) {
	db, err := service.db.GetDatabase()
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	account := &model.User{}
	err = db.C("accounts").FindId(bson.ObjectIdHex(userId)).One(&account)
	if err == nil {
		return nil, common.NewServiceError(http.StatusConflict, "User already registered")
	}

	if err != mgo.ErrNotFound {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	id := bson.ObjectIdHex(userId)
	account = &model.User{
		ID: id,
		Personal: model.PersonalInformation{
			Email:     register.Email,
			BirthDate: register.Birthdate,
		},
	}
	err = db.C("accounts").Insert(account)
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	return account, nil
}
