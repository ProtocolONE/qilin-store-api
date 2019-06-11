package mapper

import (
	"encoding/json"
	"github.com/ProtocolONE/qilin-store-api/pkg/model"
	"github.com/bxcodec/faker"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserFromModel(t *testing.T) {
	shouldBe := require.New(t)

	user := model.User{}
	shouldBe.Nil(faker.FakeData(&user))
	user.ID = bson.NewObjectId()

	userDto := UserFromModel(&user)

	shouldBeEqual(t, userDto.ID, user.ID)
	shouldBeEqual(t, user.Personal, userDto.Personal)
	shouldBeEqual(t, user.Account, userDto.Account)
}

func shouldBeEqual(t *testing.T, first interface{}, second interface{}) {
	t.Helper()
	shouldBe := require.New(t)

	fJson, err := json.Marshal(first)
	shouldBe.Nil(err)

	sJson, err := json.Marshal(second)
	shouldBe.Nil(err)

	shouldBe.JSONEq(string(fJson), string(sJson))
}