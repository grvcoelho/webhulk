package server

import (
	iris "gopkg.in/kataras/iris.v8"

	cfg "github.com/grvcoelho/webhulk/pkg/config"
	mdl "github.com/grvcoelho/webhulk/pkg/middleware"
)

type Server struct {
	app  *iris.Application
	conf *cfg.Configuration
}

func (s *Server) Start() {
	s.app.Run(iris.Addr(s.conf.Server.ListenOn))
}

func New(conf *cfg.Configuration) (*Server, error) {
	logger := mdl.NewLogger()

	app := iris.New()
	app.Use(logger)

	server := &Server{app, conf}

	return server, nil
}
