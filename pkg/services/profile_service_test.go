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

type ProfileServiceTestSuite struct {
	suite.Suite
	db      *mgo.Database
	service interfaces.ProfileService
}

func Test_ProfileService(t *testing.T) {
	suite.Run(t, new(ProfileServiceTestSuite))
}

func (suite *ProfileServiceTestSuite) SetupTest() {
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

	service := &profileService{
		db: dbProvider,
	}

	suite.db = db
	suite.service = service
}

func (suite *ProfileServiceTestSuite) TearDownTest() {
	err := suite.db.DropDatabase()
	if err != nil {
		suite.Error(err)
	}
}

func (suite *ProfileServiceTestSuite) TestAccountService_GetAccount() {
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

func (suite *ProfileServiceTestSuite) TestAccountService_UpdateAccount() {
	shouldBe := require.New(suite.T())
	service := suite.service
	db := suite.db

	userId := bson.NewObjectId()
	updateDto := dto.UpdateUserDTO{
		Account: dto.UpdateAccountDTO{
			Nickname:        "test",
			PrimaryLanguage: "en",
			AdditionalLanguages: []string{
				"ru",
				"fr",
			},
		},
		Personal: dto.UpdatePersonalDTO{
			FirstName: "FirstName",
			LastName:  "LastName",
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
