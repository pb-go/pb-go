package databaseop

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoConnPool struct {
	DbCli mongo.Client
	Position int
	Available bool
}

type MongoDB struct {
	DbConn *MongoConnPool
	DbURI string
	DbColl mongo.Collection
	DefaultDB string
	DefaultColl string
	DefaultTimeout time.Time
	BsonData chan interface{}
}

func bsonM2bsonD(var1 bson.M) bson.D{

}

func bsonD2bsonM(var1 bson.D) bson.M{

}

func (mdbc MongoDB) reConn (dbCli interface{}) error {

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
