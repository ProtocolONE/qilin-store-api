package services

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"go.uber.org/zap"
	"super_api/pkg/interfaces"
)

type dbProvider struct {
	connection string
	session    *mgo.Session
	name       string
}

func NewDatabaseProvider(connectionString string, databaseName string) (interfaces.DatabaseProvider, error) {
	zap.L().Info("Creating database provider")

	bson.SetJSONTagFallback(true)
	bson.SetRespectNilValues(true)

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)

	return &dbProvider{connection: connectionString, session: session, name: databaseName}, nil
}

func (provider *dbProvider) GetDatabase() (*mgo.Database, error) {
	return provider.session.DB(provider.name), nil
}

func (dbProvider) Shutdown() {
	//TODO
}

