package main

import (
	"github.com/lavinas/keel/internal/client/adapter/hdlr/grpc"
	"github.com/lavinas/keel/pkg/config"
	"github.com/lavinas/keel/pkg/log"
	"github.com/lavinas/keel/internal/client/core/service"
	"github.com/lavinas/keel/internal/client/adapter/repo"
	"github.com/lavinas/keel/internal/client/core/domain"
)

// main is the entrypoint of the grpc application
func main() {
	l := log.NewlogFile(".", "client-grpc", true)
	c := config.NewConfig()
	r := repo.NewRepoMysql(c)
	defer r.Close()
	d := domain.NewDomain(r)
	s := service.NewService(d, c, l, r)
	grpc.StartServer(c, l, s)
}