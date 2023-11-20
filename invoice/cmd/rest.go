package main

import (
	"github.com/lavinas/keel/invoice/internal/adapter/handler"
	"github.com/lavinas/keel/invoice/internal/adapter/repository"
	"github.com/lavinas/keel/invoice/internal/adapter/tools"
	"github.com/lavinas/keel/invoice/internal/core/usecase"
)

func main() {
	config := tools.NewConfig()
	logger, err := tools.NewLogger("invoice", true)
	if err != nil {
		panic(err)
	}
	defer logger.Close()
	repo, err := repository.NewRepository(config)
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	usercase := usecase.NewUseCase(repo, logger, config)
	handler := handler.NewRest(logger, usercase)
	handler.Run()
}
