package databaseop

import (
	"context"
	"github.com/pb-go/pb-go/content_tools"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	GlobalMGC  *mongo.Client
	GlobalMDBC MongoDB
)

type MongoDB struct {
	DbConn         *mongo.Client
	DbURI          string
	DbColl         mongo.Collection
	DefaultDB      string
	DefaultColl    string
	DefaultTimeout time.Time
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

func (mdbc *MongoDB) InitMDBCOptions() *options.ClientOptions {
	clientOptions := options.Client()
	clientOptions.ApplyURI(mdbc.DbURI)
	clientOptions.SetMinPoolSize(2)
	clientOptions.SetMaxPoolSize(4)
	clientOptions.SetRetryReads(true)
	clientOptions.SetRetryWrites(true)
	clientOptions.SetConnectTimeout(5 * time.Second)
	clientOptions.SetSocketTimeout(8 * time.Second)
	return clientOptions
}

func (mdbc *MongoDB) ConnNCheck(dbCliOption interface{}) error {
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

func (mdbc MongoDB) ItemCreate(inputdata interface{}) error {
	if inputdata == nil {
		return errors.New("Insert Queue Empty.")
	} else {
		tctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		tempData := inputdata
		insertRes, err := mdbc.DbColl.InsertOne(tctx, tempData)
		if insertRes != nil && err == nil {
			log.Println("DB Inserted a single document: ", insertRes.InsertedID)
		}
		return err
	}
}

func (mdbc MongoDB) ItemRead(filter1 interface{}) (UserData, error) {
	tctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if (mdbc.DbColl == mongo.Collection{}) {
		return UserData{}, errors.New("Default connection to coll is not setup.")
	}
	var queryRes UserData
	err := mdbc.DbColl.FindOne(tctx, filter1).Decode(&queryRes)
	if err != nil && !queryRes.EqualsTo(UserData{}){
		return UserData{}, err
	} else {
		return queryRes, nil
	}
}

func (mdbc MongoDB) ItemUpdate(filter1 interface{}, change1 interface{}) error {
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

func (mdbc MongoDB) ItemDelete(filter1 interface{}) error {
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

func (dt UserData) EqualsTo(comparedto UserData) bool {
	var check5 bool = false
	check1 := dt.WaitVerify == comparedto.WaitVerify
	check2 := dt.Data.Subtype == dt.Data.Subtype
	check3 := len(dt.Data.Data) == len(dt.Data.Data)
	if check2 && check3 {
		tmpvar1 := content_tools.GenBlake2B(dt.Data.Data)
		tmpvar2 := content_tools.GenBlake2B(comparedto.Data.Data)
		if tmpvar1 == tmpvar2 {
			check5 = true
		}
	} else {
		return false
	}
	check4 := dt.ExpireAt == comparedto.ExpireAt
	check6 := dt.Password == comparedto.Password
	check7 := dt.PwdIsSet == comparedto.PwdIsSet
	check8 := dt.ShortId == comparedto.ShortId
	return check1 && check2 && check3 && check4 && check5 && check6 && check7 && check8
}
