package main

import (
	"github.com/lavinas/keel/internal/client/core/service"
	"github.com/lavinas/keel/internal/client/adapter/hdlr"
	"github.com/lavinas/keel/internal/client/adapter/repo"
	"github.com/lavinas/keel/pkg/config"
	"github.com/lavinas/keel/pkg/log"
)

// main is the entrypoint of the application
func main() {
	c := config.NewConfig()
	l := log.NewlogFile(".", "client", true)
	r := repo.NewRepoMysql(c)
	s := service.NewService(l, r)
	h := hdlr.NewHandlerGin(l, s)
	h.Run()
}
