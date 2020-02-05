package databaseop

import (
	"github.com/pb-go/pb-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
	"time"
)

var globalMGC *mongo.Client

func TestMongoDBConn(t *testing.T) {
	var mgcli = MongoDB{
		DbConn:         globalMGC,
		DbURI:          "mongodb://localhost:27017",
		DbColl:         mongo.Collection{},
		DefaultDB:      "pbgo",
		DefaultColl:    "userdata",
		DefaultTimeout: time.Time{},
	}
	clientOptions := options.Client()
	clientOptions.ApplyURI(mgcli.DbURI)
	clientOptions.SetMinPoolSize(2)
	clientOptions.SetMaxPoolSize(4)
	clientOptions.SetRetryReads(true)
	clientOptions.SetRetryWrites(true)
	clientOptions.SetConnectTimeout(5 * time.Second)
	clientOptions.SetSocketTimeout(8 * time.Second)
	err := mgcli.ConnNCheck(clientOptions)
	mgcli.DbColl = *mgcli.DbConn.Database(mgcli.DefaultDB).Collection(mgcli.DefaultColl)
	if err != nil {
		t.Fail()
	}
	var tempIP string
	tempIP, err = utils.IP2Intstr("113.55.13.1")
	if err != nil {
		t.Fail()
	}
	var IPval primitive.Decimal128
	IPval, _ = primitive.ParseDecimal128(tempIP)
	var UserDT primitive.DateTime
	UserDT = primitive.NewDateTimeFromTime(time.Now().Add(24 * time.Hour))
	testdt1 := UserData{
		WaitVerify:   true,
		ReadThenBurn: true,
		ShortId:      "2s4D",
		UserIP:       IPval,
		ExpireAt:     UserDT,
		Data:         utils.Pack2BinData([]byte("testdata001")),
		PwdIsSet:     true,
		Password:     "He1loWorld234",
	}
	err = mgcli.ItemCreate(testdt1)
	if err != nil {
		log.Println("Failed to create document")
		t.Fail()
	}
	filter1 := bson.M{"shortId": "2s4D"}
	var readOutData UserData
	readOutData, err = mgcli.ItemRead(filter1)
	if err != nil && readOutData.EqualsTo(UserData{})  {
		t.Fail()
	} else {
		log.Println(readOutData)
	}
	time.Sleep(5 * time.Second)
	update1 := bson.D{
		{"$set", bson.D{
			{"data", utils.Pack2BinData("testdata002")},
		}},
	}
	err = mgcli.ItemUpdate(filter1, update1)
	if err != nil {
		t.Fail()
	}
	time.Sleep(5 * time.Second)
	err = mgcli.ItemDelete(filter1)
	if err != nil {
		t.Fail()
	}
	log.Println("Test Done!")
	os.Exit(0)
}
