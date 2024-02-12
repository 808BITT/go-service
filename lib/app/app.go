package app

import (
	"os"
	"os/signal"
	"sampleservice/lib/logger"
	"sync"
	"syscall"
	"time"
)

func Run(stopFlag *bool, wg *sync.WaitGroup) {
	log := logger.NewLogger("C:/dev/bobbitt/go-service/logs", true)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		defer wg.Done()
		for {
			if *stopFlag {
				log.Info("Stopping")
				// Add cleanup code if needed
				log.Info("Stopped")
				break
			}
			log.Info("Running")
			time.Sleep(5 * time.Second)
		}
	}()

	<-signalChan
	*stopFlag = true
}
