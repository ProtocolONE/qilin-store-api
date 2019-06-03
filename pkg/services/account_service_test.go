package services

import (
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/common"
	"github.com/ProtocolONE/qilin-store-api/pkg/conf"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountServiceTestSuite struct {
	suite.Suite
	db      *mgo.Database
	service interfaces.AccountService
}

func Test_AccountService(t *testing.T) {
	suite.Run(t, new(AccountServiceTestSuite))
}

func (suite *AccountServiceTestSuite) SetupTest() {
	cfg, err := conf.GetConfig()
	if err != nil {
		suite.FailNow("Config load failed", err.Error())
	}

	dbProvider, err := NewDatabaseProvider(cfg.Db)
	if err != nil {
		suite.FailNow("Can't get db provider", err.Error())
	}

	db, err := dbProvider.GetDatabase()

	if err != nil {
		suite.FailNow("Can't get database", err.Error())
	}

	service := &accountService{
		db: dbProvider,
	}

	suite.db = db
	suite.service = service
}

func (suite *AccountServiceTestSuite) TearDownTest() {
	err := suite.db.DropDatabase()
	if err != nil {
		suite.Error(err)
	}
}

func (suite *AccountServiceTestSuite) TestAccountService_Register() {
	shouldBe := require.New(suite.T())
	service := suite.service
	db := suite.db
	user, err := service.Register(bson.NewObjectId().Hex(), dto.RegisterAccountDTO{Email: "test@email.com"})
	shouldBe.Nil(err)
	shouldBe.NotNil(user)
	shouldBe.Equal("test@email.com", user.Personal.Email)

	userId := bson.NewObjectId()
	shouldBe.Nil(db.C("accounts").Insert(&model.User{ID: userId}))
	user, err = service.Register(userId.Hex(), dto.RegisterAccountDTO{})
	shouldBe.Nil(user)
	shouldBe.NotNil(err)
	shouldBe.Equal(409, err.(*common.ServiceError).Code)
}

func (suite *AccountServiceTestSuite) TestAccountService_Authorize() {
	shouldBe := require.New(suite.T())
	service := suite.service
	db := suite.db

	user, err := service.Authorize(bson.NewObjectId().Hex(), dto.AuthorizeAccountDTO{})
	shouldBe.Nil(user)
	shouldBe.NotNil(err)
	shouldBe.Equal(404, err.(*common.ServiceError).Code)

	userId := bson.NewObjectId()
	shouldBe.Nil(db.C("accounts").Insert(&model.User{ID: userId}))
	user, err = service.Authorize(userId.Hex(), dto.AuthorizeAccountDTO{})
	shouldBe.Nil(err)
	shouldBe.NotNil(user)
}

func (suite *AccountServiceTestSuite) TestAccountService_GetAccount() {
	shouldBe := require.New(suite.T())
	service := suite.service
	db := suite.db

	userId := bson.NewObjectId()
	user, err := service.GetAccount(userId.Hex())
	shouldBe.Nil(user)
	shouldBe.NotNil(err)
	shouldBe.Equal(404, err.(*common.ServiceError).Code)

	shouldBe.Nil(db.C("accounts").Insert(&model.User{ID: userId}))
	user, err = service.GetAccount(userId.Hex())
	shouldBe.Nil(err)
	shouldBe.NotNil(user)
}

func (suite *AccountServiceTestSuite) TestAccountService_RemoveMFA() {
	shouldBe := require.New(suite.T())
	service := suite.service
	db := suite.db

	userId := bson.NewObjectId()
	providerId := "111222333"
	user, err := service.RemoveMFA(userId.Hex(), providerId)
	shouldBe.Nil(user)
	shouldBe.NotNil(err)
	shouldBe.Equal(404, err.(*common.ServiceError).Code)

	insertUser := &model.User{
		ID: userId,
		Security: &model.UserSecurity{
			MFA: []model.UserMFA{
				{
					ProviderName: providerId,
					ProviderId: providerId,
				},
			},
		},
	}
	shouldBe.Nil(db.C("accounts").Insert(insertUser))

	user, err = service.RemoveMFA(userId.Hex(), providerId)
	shouldBe.Nil(err)
	shouldBe.NotNil(user)

	user, err = service.RemoveMFA(userId.Hex(), providerId)
	shouldBe.Nil(user)
	shouldBe.NotNil(err)
	shouldBe.Equal(400, err.(*common.ServiceError).Code)
}

func (suite *AccountServiceTestSuite) TestAccountService_UpdateAccount() {
	shouldBe := require.New(suite.T())
	service := suite.service
	db := suite.db

	userId := bson.NewObjectId()
	updateDto := dto.UpdateUserDTO{
		Account: dto.UpdateAccountDTO{
			Nickname: "test",
			PrimaryLanguage: "en",
			AdditionalLanguages: []string{
				"ru",
				"fr",
			},
		},
		Personal: dto.UpdatePersonalDTO{
			FirstName: "FirstName",
			LastName: "LastName",
		},
	}
	user, err := service.UpdateAccount(userId.Hex(), updateDto)
	shouldBe.Nil(user)
	shouldBe.NotNil(err)
	shouldBe.Equal(404, err.(*common.ServiceError).Code)

	insertUser := &model.User{
		ID: userId,
	}
	shouldBe.Nil(db.C("accounts").Insert(insertUser))

	user, err = service.UpdateAccount(userId.Hex(), updateDto)
	shouldBe.Nil(err)
	shouldBe.NotNil(user)
}

func (suite *AccountServiceTestSuite) TestAccountService_AddMFA() {
	shouldBe := require.New(suite.T())
	service := suite.service
	db := suite.db

	userId := bson.NewObjectId()
	providerId := "111222333"
	user, err := service.AddMFA(userId.Hex(), providerId, providerId)
	shouldBe.Nil(user)
	shouldBe.NotNil(err)
	shouldBe.Equal(404, err.(*common.ServiceError).Code)

	insertUser := &model.User{
		ID: userId,
	}
	shouldBe.Nil(db.C("accounts").Insert(insertUser))

	user, err = service.AddMFA(userId.Hex(), providerId, providerId)
	shouldBe.Nil(err)
	shouldBe.NotNil(user)

	user, err = service.AddMFA(userId.Hex(), providerId, providerId)
	shouldBe.Nil(user)
	shouldBe.NotNil(err)
	shouldBe.Equal(409, err.(*common.ServiceError).Code)
}