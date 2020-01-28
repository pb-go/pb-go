package databaseop

import (
	"github.com/matoous/go-nanoid"
	"log"
)

type DBClient interface {
	connect2DB(dbURI string) error
	checkConn(dbCli interface{}) error
	testCollectionNIndex(dbCli interface{}, dbName string, collName string) int
	itemCreate(dbCli interface{}, dbName string, collName string, data interface{}) int
	itemUpdate(dbCli interface{}, dbName string, collName string, data interface{}) int
	itemDelete(dbCli interface{}, dbName string, collName string, data interface{}) int
	itemRead(dbCli interface{}, dbName string, collName string, data interface{}) int
}

func getNanoID() (string, error) {
	id, err := gonanoid.Nanoid(4)
	if err != nil {
		log.Fatalln("Failed to generate nanoid!")
	}
	return id,err
}