package main

import (
	"os"
	"sampleservice/lib/app"
	"sampleservice/lib/config"
	"sampleservice/lib/logger"
	"sampleservice/lib/service"
	"sync"

	"golang.org/x/sys/windows/svc"
)

var log *logger.Logger
var configuration *config.Config
var wg sync.WaitGroup
var stopFlag = new(bool)

func init() {
	if len(os.Args) != 2 {
		log.Error("Usage: .../main.exe <config path>")
		return
	}
	configuration = config.NewConfig(os.Args[1])
	if configuration == nil {
		os.Exit(1)
	}

	log = logger.NewLogger(configuration.Install.Path+"/logs", true)
}

func main() {
	isWinServ, err := svc.IsWindowsService()
	if err != nil {
		log.Error("failed to determine if we are running as a windows service: " + err.Error())
		return
	}
	if isWinServ {
		serviceName := configuration.Install.Name
		service.Run(serviceName)
		return
	}

	wg.Add(1)
	go func() {
		app.Run(configuration, log, stopFlag, &wg)
	}()

	wg.Wait()
	log.Info("Exited")
}
