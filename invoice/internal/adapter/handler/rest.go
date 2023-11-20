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
func NewRest(logger port.Logger, usercase port.UseCase) *Rest {
	return &Rest{logger: logger, usercase: usercase, krest: krest.NewKrest()}
}

// Run runs the rest handler
func (h *Rest) Run() {
	handlers := krest.HandlerMap{
		"GET":  {"/invoice/ping": h.ping},
		"POST": {"/invoice/client": h.registerClient},
	}
	h.krest.Run(h.logger, handlers)
}

// ping is the ping handler
func (h *Rest) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (h *Rest) registerClient(c *gin.Context) {
	var input dto.RegisterClient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usercase.RegisterClient(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
