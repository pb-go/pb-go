package main

import (
	"fmt"
	"github.com/pb-go/pb-go/clipkg"
	"github.com/pb-go/pb-go/config"
	"log"
	"os"
)

// main function for client
func main() {
	fmt.Fprintln(os.Stderr, "Version: "+config.CurrentVer)
	err := clipkg.Execute()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
