package main

import (
	"github.com/pb-go/pb-go/pkg/command"
	"log"
)

func main() {
	err := command.Execute()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
