package service

import (
	"fmt"
	"os"
	"sampleservice/lib/app"
	"sampleservice/lib/config"
	"sampleservice/lib/logger"
	"strings"
	"sync"

	"golang.org/x/sys/windows/svc"
)

var configuration *config.Config
var log *logger.Logger
var wg sync.WaitGroup

type Service struct {
	StopFlag bool
}

func init() {
	configPath := os.Args[1]
	fmt.Println("configPath: ", configPath)
	configuration = config.NewConfig(configPath)
	if configuration == nil {
		os.Exit(1)
	}

	log = logger.NewLogger(configuration.Install.Path+"/logs", true)
}

func (s *Service) Execute(_ []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}

	wg.Add(1)
	go func() {
		app.Run(configuration, log, &s.StopFlag, &wg)
	}()

	log.Info("Service started")

	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	for c := range r {
		switch c.Cmd {
		case svc.Stop, svc.Shutdown:
			s.StopFlag = true
			changes <- svc.Status{State: svc.StopPending}
			wg.Wait()
			return
		default:
			log.Warning("Command not recognized: " + fmt.Sprint(c.Cmd))
		}
	}
	return
}

func Run(name string) {
	log := logger.NewLogger(strings.Split(os.Args[0], "main.exe")[0]+"logs", true)
	log.Info("Service starting")
	err := svc.Run(name, &Service{})
	if err != nil {
		log.Error("Service failed: " + err.Error())
	}
	log.Info("Service exited.")
}
