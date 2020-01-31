package databaseop

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoConn struct {
	DbCli mongo.Client
	Position int
	Available bool
}

type MongoDB struct {
	DbConn         *mongo.Client
	DbURI          string
	DbColl         mongo.Collection
	DefaultDB      string
	DefaultColl    string
	DefaultTimeout time.Time
	BsonWData      chan UserData
	BsonRData      chan UserData
}

type UserData struct {
	WaitVerify bool                 `bson:"waitVerify" json:"waitVerify"`
	ShortId    string               `bson:"shortId" json:"shortId"`
	UserIP     primitive.Decimal128 `bson:"userIP" json:"userIP"`
	ExpireAt   primitive.DateTime   `bson:"expireAt" json:"expireAt"`
	Data       primitive.Binary     `bson:"data" json:"data"`
	PwdIsSet   bool                 `bson:"pwdIsSet" json:"pwdIsSet"`
	Password   string               `bson:"passwd" json:"passwd"`
}

// only allow bson.M to be used

func (mdbc *MongoDB) connNCheck(dbCliOption interface{}) error {
	// https://github.com/mongodb/mongo-go-driver/blob/master/data/connection-monitoring-and-pooling/connection-monitoring-and-pooling.rst
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	var err error
	mdbc.DbConn, err = mongo.Connect(ctx, dbCliOption.(*options.ClientOptions))
	// URI with srv must not include a port number
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot connect to DB.")
		return err
	}
	log.Println("Database Connection Get, Testing...")
	err = mdbc.DbConn.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("DB Connection is not responding.")
		return err
	} else {
		log.Println("Database Successfully Connected!")
		return nil
	}
}

func (mdbc MongoDB) itemCreate() error {
	if len(mdbc.BsonWData) == 0 {
		return errors.New("Insert Queue Empty.")
	} else {
		tctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		tempData := <-mdbc.BsonWData
		insertRes, err := mdbc.DbColl.InsertOne(tctx, tempData)
		if insertRes != nil && err == nil {
			log.Println("DB Inserted a single document: ", insertRes.InsertedID)
		}
		return err
	}
}

func (mdbc MongoDB) itemRead(filter1 interface{}) error {
	tctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if (mdbc.DbColl == mongo.Collection{}) {
		return errors.New("Default connection to coll is not setup.")
	}
	var queryRes UserData
	err := mdbc.DbColl.FindOne(tctx, filter1).Decode(&queryRes)
	if err != nil {
		return err
	} else {
		mdbc.BsonRData <- queryRes
		return nil
	}
}

func (mdbc MongoDB) itemUpdate(filter1 interface{}, change1 interface{}) error {
	tctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if (mdbc.DbColl == mongo.Collection{}) {
		return errors.New("Default connection to coll is not setup.")
	}
	updateRes, err := mdbc.DbColl.UpdateOne(tctx, filter1, change1)
	if err != nil {
		return err
	}
	log.Printf("Matched %v docs and updated %v docs. \n", updateRes.MatchedCount, updateRes.ModifiedCount)
	return nil
}

func (mdbc MongoDB) itemDelete(filter1 interface{}) error {
	if (mdbc.DbColl == mongo.Collection{}) {
		return errors.New("Connection to coll is not setup.")
	}
	tctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	deleteRes, err := mdbc.DbColl.DeleteOne(tctx, filter1)
	if err != nil {
		return err
	}
	log.Printf("Deleted %v documents.", deleteRes.DeletedCount)
	return nil
}

