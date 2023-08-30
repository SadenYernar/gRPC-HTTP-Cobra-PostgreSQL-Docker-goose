package api

import (
	"net/http"
	"time"

	"github.com/SadenYernar/gRPC-HTTP-Cobra-PostgreSQL-Docker-goose.git/internal/config"
	"github.com/SadenYernar/gRPC-HTTP-Cobra-PostgreSQL-Docker-goose.git/internal/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Api struct {
	c *config.Config
	// s *server.Server
	m  *runtime.ServeMux
	gs *grpc.Server
	hs *http.Server
}

func Get(c *config.Config /*s *server.Server*/, m *runtime.ServeMux) *Api {
	return &Api{
		c: c,

		m:  m,
		gs: grpc.NewServer(),
		hs: &http.Server{
			Addr:         c.HttpPort,
			ErrorLog:     logger.ErrorImpl,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}
}
