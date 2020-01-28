package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "github.com/fvbock/endless"
	"fmt"
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var (
	currentVer string = "v0.1.0"
	version = flag.Bool("version", false, "Show current version of pb-go.")
	confFile = flag.String("config", "config.yaml", "Server config for pb-go.")
	servConf ServConfig
)

func printVersion() {
	fmt.Printf("Current Version: %s \n", currentVer)
	fmt.Println("For more information: https://github.com/kmahyyg/pb-go")
	fmt.Println("This Program is licensed under AGPLv3.")
}

func fileExist(filepath string) bool{
	info, err := os.Stat(filepath)
	return err == nil && !info.IsDir()
}

func startServer(conf ServConfig) error{

}

func main() {
	flag.Parse()

	if *version {
		printVersion()
	}

	if workingDir, err := os.Getwd(); err == nil {
		confPath := filepath.Join(workingDir, *confFile)
		if fileExist(confPath){
			servConf, err := loadConfig(confPath)
			if err != nil {
				os.Exit(14)
			} else {
				err := startServer(servConf)
				os.Exit(0)
			}
		}
	} else {
		os.Exit(13)
	}

	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-osSignals
	}


}
