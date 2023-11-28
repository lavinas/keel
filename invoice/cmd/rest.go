package main

import (
	"fmt"

	"github.com/lavinas/keel/invoice/internal/adapter/handler"
	"github.com/lavinas/keel/invoice/internal/adapter/repository"
	"github.com/lavinas/keel/invoice/internal/adapter/tools"
	"github.com/lavinas/keel/invoice/internal/core/usecase"
)

// main is the main function
func main() {
	// create config
	config := tools.NewConfig()
	// create logger
	logger, err := tools.NewLogger("invoice", true)
	if err != nil {
		logger.Fatal(fmt.Errorf("fatal error creating logger: %w", err))
	}
	defer logger.Close()
	// create repository
	repo, err := repository.NewRepository(config)
	if err != nil {
		logger.Fatal(fmt.Errorf("fatal error creating repo: %w", err))
	}
	defer repo.Close()
	// create usecase
	usercase := usecase.NewUseCase(config, logger, repo)
	// create handler
	handler := handler.NewRest(config, logger, usercase)
	// run handler
	handler.Run()
}
