package app

import (
	"os"
	"os/signal"
	"sampleservice/lib/config"
	"sampleservice/lib/logger"
	"sync"
	"syscall"
	"time"
)

func Run(config *config.Config, log *logger.Logger, stopFlag *bool, wg *sync.WaitGroup) {
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGTERM, syscall.SIGINT)

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

	<-stopSignal
	*stopFlag = true
}
