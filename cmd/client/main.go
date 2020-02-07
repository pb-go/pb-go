package main

import (
	"fmt"
	"github.com/pb-go/pb-go/clipkg"
	"log"
)

const currentVersion = "v1.0.3"

// main function for client
func main() {
	fmt.Println("Version: " + currentVersion + "\n")
	err := clipkg.Execute()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
