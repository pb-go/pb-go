package main

import (
	"flag"
	"github.com/fasthttp/router"
	"github.com/getsentry/sentry-go"
	sentryfasthttp "github.com/getsentry/sentry-go/fasthttp"
	"github.com/kmahyyg/pb-go/config"
	"github.com/kmahyyg/pb-go/content_tools"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var (
	version  = flag.Bool("version", false, "Show current version of pb-go.")
	confFile = flag.String("config", "config.yaml", "Server config for pb-go.")
	fahtserv *fasthttp.Server
)

func printVersion() {
	log.Printf("Current Version: %s \n", config.CurrentVer)
	log.Println("For more information: https://github.com/kmahyyg/pb-go")
	log.Println("This Program is licensed under AGPLv3.")
}

func startServer(conf config.ServConfig) error {
	// init sentry
	var err error
	sentryHandler := sentryfasthttp.New(sentryfasthttp.Options{
		Repanic:         false,
		WaitForDelivery: false,
		Timeout:         5 * time.Second,
	})
	app := router.New()
	wrappedhand := sentryHandler.Handle(app.Handler)
	fahtserv = &fasthttp.Server{
		Handler:      wrappedhand,
		Name:         "pb-go",
		TCPKeepalive: true,
	}
	if err := fahtserv.ListenAndServe(config.ServConf.Network.Listen); err != nil {
		log.Fatalf("Error in listen on %s : %s", config.ServConf.Network.Listen, err)
	}
	return err
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: config.CurrentDSN,
	}); err != nil {
		log.Printf("Sentry Bug-Tracking init failed: %v \n", err)
	}
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
		if !filepath.IsAbs(*confFile) {
			confPath = filepath.Join(workingDir, *confFile)
		} else {
			confPath = *confFile
		}
		// check if file exists and not a directory
		if config.FileExist(confPath) {
			config.ServConf, err = config.LoadConfig(confPath)
			if err != nil {
				log.Println("Please check document on our project page.")
				os.Exit(14)
			} else {
				// start server with graceful restart
				err := startServer(config.ServConf)
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
		log.Println("Signal Received to shutdown server...")
		if err := fahtserv.Shutdown(); err != nil{
			log.Fatalf("Server Shutdown Failed: %v", err)
		}
		log.Println("Server exit successfully.")
	}

}
