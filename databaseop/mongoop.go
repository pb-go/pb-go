package databaseop

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDB struct {
	DbCli mongo.Client
	DbURI string
	DefaultDB string
	DefaultColl string
	DefaultTimeout time.Time
	BsonData chan bson.M
}

func (mdbc MongoDB) connect2DB(dbURI string) error{

}

func (mdbc MongoDB) checkConn(dbCli interface{}) error{

}

func (mdbc MongoDB) testCollectionNIndex(dbCli interface{}, dbName string, collName string) int{

}

func (mdbc MongoDB) itemCreate(dbCli interface{}, dbName string, collName string, data interface{}) int{

}

func (mdbc MongoDB) itemUpdate(dbCli interface{}, dbName string, collName string, data interface{}) int{

}

func (mdbc MongoDB) itemDelete(dbCli interface{}, dbName string, collName string, data interface{}) int{

}

func (mdbc MongoDB) itemRead(dbCli interface{}, dbName string, collName string, data interface{}) int{

}
