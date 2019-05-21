package services

import (
	"net/url"
	"time"

	"github.com/ProtocolONE/qilin-store-api/pkg/conf"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"go.uber.org/zap"
)

type dbProvider struct {
	connection string
	session    *mgo.Session
	name       string
}

func NewDatabaseProvider(c *conf.DbConfig) (interfaces.DatabaseProvider, error) {
	zap.L().Info("Creating database provider")

	bson.SetJSONTagFallback(true)
	bson.SetRespectNilValues(true)

	info, err := mgo.ParseURL(BuildConnString(c))
	if err != nil {
		return nil, err
	}

	info.Timeout = 10 * time.Second

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return nil, err
	}

	session.SetSyncTimeout(1 * time.Minute)
	session.SetSocketTimeout(1 * time.Minute)
	session.SetMode(mgo.Monotonic, true)

	return &dbProvider{connection: BuildConnString(c), session: session, name: c.Name}, nil
}

func (provider *dbProvider) GetDatabase() (*mgo.Database, error) {
	return provider.session.DB(provider.name), nil
}

func (dbProvider) Shutdown() {
	//TODO
}

func BuildConnString(c *conf.DbConfig) string {
	if c.Name == "" {
		return ""
	}

	vv := url.Values{}

	var userInfo *url.Userinfo

	if c.User != "" {
		if c.Password == "" {
			userInfo = url.User(c.User)
		} else {
			userInfo = url.UserPassword(c.User, c.Password)
		}
	}

	u := url.URL{
		Scheme:   "mongodb",
		Path:     c.Name,
		Host:     c.Host,
		User:     userInfo,
		RawQuery: vv.Encode(),
	}

	return u.String()
}
