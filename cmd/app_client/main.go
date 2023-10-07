package main

import (
	"github.com/lavinas/keel/internal/client/adapter/hdlr"
	"github.com/lavinas/keel/internal/client/adapter/repo"
	"github.com/lavinas/keel/pkg/config"
	"github.com/lavinas/keel/pkg/log"

	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/service"
)

// main is the entrypoint of the application
func main() {
	c := config.NewConfig()
	l := log.NewlogFile(".", "client", true)
	r := repo.NewRepoMysql(c)
	defer r.Close()
	d := domain.NewDomain(r)
	s := service.NewService(d, c, l, r)
	h := hdlr.NewHandlerGin(l, s)
	h.MapHandlers()
	h.Run()
}
