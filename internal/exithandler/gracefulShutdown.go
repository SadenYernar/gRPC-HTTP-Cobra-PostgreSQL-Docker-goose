package exithandler

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/SadenYernar/gRPC-HTTP-Cobra-PostgreSQL-Docker-goose.git/internal/logger"
)

func Init(cb func()) {
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logger.Info("exit reason", sig)
		terminate <- true
	}()

	<-terminate
	cb()
	logger.Info("exiting program")
}
