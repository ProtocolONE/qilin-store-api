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

const accountsCollection string = "accounts"

func NewAccountService(db interfaces.DatabaseProvider) interfaces.AccountService {
	return &accountService{db}
}

func (service *accountService) GetAccount(userId string) (*model.User, error) {
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

func (service *accountService) RemoveMFA(userId string, providerId string) (*model.User, error) {
	user, err := service.getUser(userId)
	if err != nil {
		return nil, err
	}

	if user.Security == nil {
		return nil, common.NewServiceErrorf(http.StatusBadRequest, "User don't have MFA `%s`", providerId)
	}

	var filtered []model.UserMFA
	for _, mfa := range user.Security.MFA {
		if mfa.ProviderId != providerId {
			filtered = append(filtered, mfa)
		}
	}

	if len(filtered) == len(user.Security.MFA) {
		return nil, common.NewServiceErrorf(http.StatusBadRequest, "User don't have MFA `%s`", providerId)
	}

	user.Security.MFA = filtered
	err = service.saveUser(userId, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *accountService) AddMFA(userId string, providerId string, providerName string) (*model.User, error) {
	user, err := service.getUser(userId)
	if err != nil {
		return nil, err
	}

	if user.Security == nil {
		user.Security = &model.UserSecurity{
			MFA: []model.UserMFA{},
		}
	}

	for _, mfa := range user.Security.MFA {
		if mfa.ProviderId == providerId {
			return nil, common.NewServiceErrorf(http.StatusConflict, "MFA %s already added to user", mfa.ProviderId)
		}
	}

	//TODO: здесь не хватает проверки на auth1 есть ли у пользователя МФА и верефицирована ли она. Пока что нет реализации в auth1 этого
	user.Security.MFA = append(user.Security.MFA, model.UserMFA{ProviderId: providerId, ProviderName: providerName})

	err = service.saveUser(userId, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *accountService) UpdateAccount(userId string, update dto.UpdateUserDTO) (*model.User, error) {
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

func (service *accountService) Authorize(userId string, authorize dto.AuthorizeAccountDTO) (*model.User, error) {
	db, err := service.db.GetDatabase()
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	account := &model.User{}
	err = db.C(accountsCollection).FindId(bson.ObjectIdHex(userId)).One(account)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, common.NewServiceError(http.StatusNotFound, err)
		}
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	return account, nil
}

func (service *accountService) Register(userId string, register dto.RegisterAccountDTO) (*model.User, error) {
	db, err := service.db.GetDatabase()
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	account := &model.User{}
	err = db.C(accountsCollection).FindId(bson.ObjectIdHex(userId)).One(account)
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
	err = db.C(accountsCollection).Insert(account)
	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	return account, nil
}

func (service *accountService) saveUser(userId string, user *model.User) error {
	db, err := service.db.GetDatabase()
	if err != nil {
		return common.NewServiceError(http.StatusInternalServerError, err)
	}

	err = db.C(accountsCollection).UpdateId(bson.ObjectIdHex(userId), user)
	if err != nil {
		return common.NewServiceError(http.StatusInternalServerError, err)
	}

	return nil
}

func (service *accountService) getUser(userId string) (*model.User, error) {
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

	return account, nil
}
