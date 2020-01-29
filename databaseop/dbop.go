package databaseop

import (
	"github.com/matoous/go-nanoid"
	"log"
)

type DBClient interface {
	connect2DB(dbURI string) error
	checkConn(dbCli interface{}) error
	testCollectionNIndex() int
	itemCreate() int
	itemUpdate() int
	itemDelete() int
	itemRead() int
}

func getNanoID() (string, error) {
	id, err := gonanoid.Nanoid(4)
	if err != nil {
		log.Fatalln("Failed to generate nanoid!")
	}
	return id,err
}