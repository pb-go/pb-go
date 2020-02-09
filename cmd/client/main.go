package main

import (
	"fmt"
	"github.com/pb-go/pb-go/clipkg"
	"log"
	"os"
)

const currentVersion = "v1.0.3"

// main function for client
func main() {
	fmt.Fprintln(os.Stderr, "Version: "+currentVersion)
	err := clipkg.Execute()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
