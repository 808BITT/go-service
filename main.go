package main

import (
	"os"
	"sampleservice/lib/app"
	"sampleservice/lib/logger"
	"sampleservice/lib/service"
	"sync"

	"golang.org/x/sys/windows/svc"
)

var wg sync.WaitGroup
var stopFlag = new(bool)

func main() {
	log := logger.NewLogger("C:/dev/bobbitt/go-service/logs", true)
	log.Info("Initializing")

	isWinServ, err := svc.IsWindowsService()
	if err != nil {
		log.Error("failed to determine if we are running as a windows service: " + err.Error())
		return
	}
	if isWinServ {
		if len(os.Args) < 2 {
			log.Error("Usage: main.exe <service name>")
			return
		}
		serviceName := os.Args[1]
		service.Run(serviceName)
		return
	}

	wg.Add(1)
	go func() {
		app.Run(stopFlag, &wg)
	}()

	wg.Wait()
	log.Info("Exited")
}
