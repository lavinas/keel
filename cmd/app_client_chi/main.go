package main

import (
	"github.com/lavinas/keel/internal/client/core/service"
	"github.com/lavinas/keel/internal/client/hdlr"
	"github.com/lavinas/keel/internal/client/repo"
	"github.com/lavinas/keel/pkg/config"
)

func main() {
	config := config.NewConfig()
	repo := repo.NewRepoMysql(config)
	defer repo.Close()
	service := service.NewService(repo)
	hdlr.NewHandlerChi(service)
}