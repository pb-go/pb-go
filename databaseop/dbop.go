package databaseop

import (
	"github.com/matoous/go-nanoid"
	"log"
	"os"
)

type DBClient interface {

}

func getNanoID() (string, error) {
	id, err := gonanoid.Nanoid(4)
	if err != nil {
		log.Fatalln("Failed to generate nanoid!")
	}
	return id,err
}