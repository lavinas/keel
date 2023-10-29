package main

import (
	"github.com/lavinas/keel/internal/invoice/adapter/hdlr/rest"
	"github.com/lavinas/keel/internal/invoice/adapter/repo/mysql"
	"github.com/lavinas/keel/pkg/log"

	"github.com/lavinas/keel/internal/invoice/core/domain"
	"github.com/lavinas/keel/internal/invoice/core/service"
)

// main is the entrypoint of the rest application
func main() {
	l, err := log.NewlogFile("client-rest", true)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	r, err := mysql.NewRepoMysql()
	if err != nil {
		panic(err)
	}
	defer r.Close()
	d := domain.NewDomain(r)
	s := service.NewService(l, d)
	h := rest.NewHandlerRest(l, s)
	h.Run()
}
