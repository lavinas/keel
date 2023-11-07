package main

import (
	"github.com/lavinas/keel/client/internal/adapter/hdlr/rest"
	"github.com/lavinas/keel/client/internal/adapter/hdlr/config"
	"github.com/lavinas/keel/client/internal/adapter/repo/mysql"
	"github.com/lavinas/keel/client/pkg/log"

	"github.com/lavinas/keel/client/internal/core/domain"
	"github.com/lavinas/keel/client/internal/core/service"
)

// main is the entrypoint of the rest application
func main() {
	g := config.NewConfig()
	l, err := log.NewlogFile("client-rest", true)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	r := mysql.NewRepoMysql(g)
	defer r.Close()
	d := domain.NewDomain(r)
	s := service.NewService(d, g, l)
	h := rest.NewHandlerRest(l, s)
	h.Run()
}
