package main

import (
	"flag"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	log.Printf("Current Version: %s \n", currentVer)
	log.Println("For more information: https://github.com/kmahyyg/pb-go")
	log.Println("This Program is licensed under AGPLv3.")
}

func fileExist(filepath string) bool{
	info, err := os.Stat(filepath)
	return err == nil && !info.IsDir()
}

func startServer(conf ServConfig) error{
	//todo: graceful restart
	//todo: custom port and listen host
	//todo: gin framework
}

func init(){
	log.SetFlags(log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
}

func main() {
	flag.Parse()

	printVersion()

	// after parsing command line args, do corresponding operation
	if *version {
		os.Exit(0)
	}

	if workingDir, err := os.Getwd(); err == nil {
		var confPath string
		// if user doesn't offer absolute path of config file
		if !filepath.IsAbs(*confFile){
			confPath = filepath.Join(workingDir, *confFile)
		} else {
			confPath = *confFile
		}
		// check if file exists and not a directory
		if fileExist(confPath){
			servConf, err := loadConfig(confPath)
			if err != nil {
				log.Println("Please check document on our project page.")
				os.Exit(14)
			} else {
				// start server with graceful restart
				err := startServer(servConf)
				if err != nil {
					os.Exit(1)
				}
				defer os.Exit(0)
			}
		}
	} else {
		os.Exit(13)
	}

	// handler of user issued system signal
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-osSignals
	}

}
