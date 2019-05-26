package services

import (
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/common"
	"github.com/ProtocolONE/qilin-store-api/pkg/conf"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccountService_Register(t *testing.T) {
	shouldBe := require.New(t)
	cfg, err := conf.GetConfig()
	if err != nil {
		panic(err)
	}

	dbProvider, err := NewDatabaseProvider(cfg.Db)

	service := &accountService{
		db: dbProvider,
	}

	user, err := service.Register(bson.NewObjectId().Hex(), &dto.RegisterAccountDTO{Email: "test@email.com"})
	shouldBe.Nil(err)
	shouldBe.NotNil(user)
	shouldBe.Equal("test@email.com", user.Personal.Email)

	db, err := dbProvider.GetDatabase()
	shouldBe.Nil(err)
	userId := bson.NewObjectId()
	shouldBe.Nil(db.C("accounts").Insert(&model.User{ID: userId}))
	user, err = service.Register(userId.Hex(), &dto.RegisterAccountDTO{})
	shouldBe.Nil(user)
	shouldBe.NotNil(err)
	shouldBe.Equal(409, err.(*common.ServiceError).Code)
}

func TestAccountService_Authorize(t *testing.T) {
	shouldBe := require.New(t)
	cfg, err := conf.GetConfig()
	if err != nil {
		panic(err)
	}

	dbProvider, err := NewDatabaseProvider(cfg.Db)

	service := &accountService{
		db: dbProvider,
	}

	user, err := service.Authorize(bson.NewObjectId().Hex(), &dto.AuthorizeAccountDTO{})
	shouldBe.Nil(user)
	shouldBe.NotNil(err)
	shouldBe.Equal(404, err.(*common.ServiceError).Code)

	db, err := dbProvider.GetDatabase()
	shouldBe.Nil(err)
	userId := bson.NewObjectId()
	shouldBe.Nil(db.C("accounts").Insert(&model.User{ID: userId}))
	user, err = service.Authorize(userId.Hex(), &dto.AuthorizeAccountDTO{})
	shouldBe.Nil(err)
	shouldBe.NotNil(user)
}


