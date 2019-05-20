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

	userJson, err := json.Marshal(user)
	shouldBe.Nil(err)
	dtoJson, err := json.Marshal(userDto)
	shouldBe.Nil(err)

	shouldBe.JSONEq(string(userJson), string(dtoJson))
}
