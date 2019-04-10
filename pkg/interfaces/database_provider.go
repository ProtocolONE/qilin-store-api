package interfaces

import "github.com/globalsign/mgo"

type DatabaseProvider interface {
	GetDatabase() (*mgo.Database, error)
	Shutdown()
}
