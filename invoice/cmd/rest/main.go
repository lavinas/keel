package main

import (
	"github.com/lavinas/keel/invoice/internal/adapter/hdlr/config"
	"github.com/lavinas/keel/invoice/internal/adapter/hdlr/rest"
	"github.com/lavinas/keel/invoice/internal/adapter/hdlr/rest_consumer"
	"github.com/lavinas/keel/invoice/internal/adapter/repo/mysql"
	"github.com/lavinas/keel/invoice/pkg/log"

	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/service"
)

// main is the entrypoint of the rest application
func main() {
	g := config.NewConfig()
	l, err := log.NewlogFile("invoice-rest", true)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	r, err := mysql.NewRepoMysql(g)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	c := restconsumer.NewRestConsumer(g)
	d := domain.NewDomain(r)
	s := service.NewService(l, c, d)
	h := rest.NewHandlerRest(l, s)
	h.Run()
}
