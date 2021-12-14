package confhttp

import (
	"context"
	"fmt"

	"github.com/go-jarvis/rum-gonic/rum"
)

type Server struct {
	appname string
	Addr    string `env:""`
	Port    int    `env:""`
	*rum.Engine
}

/* jarvis config adpat */

func (s *Server) SetDefaults() {
	if s.Port == 0 {
		s.Port = 80
	}

	if s.appname == "" {
		s.appname = "rum-gonic httpserver"
	}
}

func (s *Server) initial() {
	if s.Engine == nil {
		s.Engine = rum.Default()
	}
}

func (s *Server) Init() {
	s.SetDefaults()
	s.initial()
}

func (s *Server) SetAppname(app string) {
	s.appname = fmt.Sprintf("%s %s", app, "rum-httpserver")
}

/* jarvis launcher adapt*/

func (s *Server) Appname() string {
	return s.appname
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", s.Addr, s.Port)
	return s.Engine.ListenAndServe(addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Engine.Shutdown(ctx)
}
