package main

import (
	"fmt"

	"github.com/lavinas/keel/internal/invoice/adapter/handler"
	"github.com/lavinas/keel/internal/invoice/adapter/repository"
	"github.com/lavinas/keel/internal/invoice/adapter/tools"
	"github.com/lavinas/keel/internal/invoice/core/usecase"
)

const (
	name                     = "invoice"
	fatalErrorCreatingLogger = "fatal error creating logger: %w"
	fatalErrorCreatingRepo   = "fatal error creating repo: %w"
)

// main is the main function
func main() {
	// create config
	config := tools.NewConfig()
	// create logger
	logger, err := tools.NewLogger(name, true)
	if err != nil {
		logger.Fatal(fmt.Errorf(fatalErrorCreatingLogger, err))
	}
	defer logger.Close()
	// create repository
	repo, err := repository.NewRepository(config)
	if err != nil {
		logger.Fatal(fmt.Errorf(fatalErrorCreatingRepo, err))
	}
	defer repo.Close()
	// create usecase
	usercase := usecase.NewUseCase(config, logger, repo)
	// create handler
	handler := handler.NewRest(config, logger, usercase)
	// run handler
	handler.Run()
}
