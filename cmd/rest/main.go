package main

import (
	"github.com/lavinas/keel/internal/client/adapter/hdlr/rest"
	"github.com/lavinas/keel/internal/client/adapter/repo/mysql"
	"github.com/lavinas/keel/pkg/log"

	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/service"
)

// main is the entrypoint of the rest application
func main() {
	l, err := log.NewlogFile("client-rest", true)
	if err != nil {
		panic(err)
	}
	r := mysql.NewRepoMysql()
	defer r.Close()
	d := domain.NewDomain(r)
	s := service.NewService(d, l, r)
	h := rest.NewHandlerRest(l, s)
	h.Run()
}
