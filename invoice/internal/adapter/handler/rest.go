package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/invoice/internal/core/dto"
	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/krest"
)

// Rest is the rest handler for the application
type Rest struct {
	logger   port.Logger
	usercase port.UseCase
	krest    *krest.Krest
}

// NewRest creates a new rest handler
func NewRest(config port.Config, logger port.Logger, usercase port.UseCase) *Rest {
	return &Rest{logger: logger, usercase: usercase, krest: krest.NewKrest(config, logger)}
}

// Run runs the rest handler
func (h *Rest) Run() {
	handlers := krest.HandlerMap{
		"GET": {
			"/invoice/ping": h.ping,
		},
		"POST": {
			"/invoice/client":      h.Register,
			"/invoice/instruction": h.Register,
			"/invoice/product":     h.Register,
		},
	}
	h.krest.Run(handlers)
}

// ping is the ping handler
func (h *Rest) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// registerClient is the register client handler
func (h *Rest) Register(c *gin.Context) {
	input := h.registerFactory(c)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 409, "message": "invalid json structure"})
		return
	}
	var result dto.DefaultResult
	h.usercase.Register(input, &result)
	c.JSON(result.Code, result)
}

func (h *Rest) registerFactory(c *gin.Context) port.Register {
	switch c.Request.URL.Path {
	case "/invoice/client":
		return &dto.RegisterClient{}
	case "/invoice/instruction":
		return &dto.RegisterInstruction{}
	case "/invoice/product":
		return &dto.RegisterProduct{}
	}
	return nil
}