package main

import (
	"github.com/lavinas/keel/invoice/internal/adapter/hdlr/rest"
	"github.com/lavinas/keel/invoice/internal/adapter/repo/mysql"
	"github.com/lavinas/keel/util/pkg/log"

	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/service"
)

// main is the entrypoint of the rest application
func main() {
	l, err := log.NewlogFile("invoice-rest", true)
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
