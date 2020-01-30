package databaseop

import (
	"github.com/matoous/go-nanoid"
	"log"
)

type DBClient interface {
	connNCheck(dbCliOption interface{}) error
	itemCreate() error
	itemUpdate(filter1 interface{}, change1 interface{}) error
	itemDelete(filter1 interface{}) error
	itemRead(filter1 interface{}) error
}

func getNanoID() (string, error) {
	id, err := gonanoid.Nanoid(4)
	if err != nil {
		log.Fatalln("Failed to generate nanoid!")
	}
	return id,err
}