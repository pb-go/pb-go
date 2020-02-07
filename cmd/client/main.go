package main

import (
	"github.com/pb-go/pb-go/clipkg"
	"log"
)

// main function for client
func main() {
	err := clipkg.Execute()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
