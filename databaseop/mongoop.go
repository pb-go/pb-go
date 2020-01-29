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

func bsonM2bsonD(var1 bson.M) bson.D{

}

func bsonD2bsonM(var1 bson.D) bson.M{

}

func (mdbc MongoDB) connect2DB(dbURI string) error{

}

func (mdbc MongoDB) checkConn(dbCli interface{}) error{

}

func (mdbc MongoDB) testCollectionNIndex() int{

}

func (mdbc MongoDB) itemCreate() int{

}

func (mdbc MongoDB) itemUpdate() int{

}

func (mdbc MongoDB) itemDelete() int{

}

func (mdbc MongoDB) itemRead() int{

}
