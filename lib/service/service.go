package service

import (
	"fmt"
	"sampleservice/lib/app"
	"sampleservice/lib/logger"
	"sync"

	"golang.org/x/sys/windows/svc"
)

var wg sync.WaitGroup

type Service struct {
	StopFlag bool
}

func (s *Service) Execute(_ []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	log := logger.NewLogger("C:/dev/bobbitt/go-service/logs", true)

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}

	wg.Add(1)
	go func() {
		app.Run(&s.StopFlag, &wg)
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
	log := logger.NewLogger("C:/dev/bobbitt/go-service/logs", true)
	log.Info("Service starting")
	err := svc.Run(name, &Service{})
	if err != nil {
		log.Error("Service failed: " + err.Error())
	}
	log.Info("Service exited.")
}
