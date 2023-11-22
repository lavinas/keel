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
			"/invoice/client":      h.registerClient,
			"/invoice/instruction": h.registerInstruction,
			"/invoice/product":     h.registerProduct,
		},
	}
	h.krest.Run(handlers)
}

// ping is the ping handler
func (h *Rest) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// registerClient is the register client handler
func (h *Rest) registerClient(c *gin.Context) {
	var input dto.RegisterClient
	if !h.bindInput(c, &input) {
		return
	}
	var result dto.DefaultResult
	h.usercase.Register(&input, &result)
	c.JSON(result.Code, result)
}

// registerInstruction is the register instruction handler
func (h *Rest) registerInstruction(c *gin.Context) {
	var input dto.RegisterInstruction
	if !h.bindInput(c, &input) {
		return
	}
	var result dto.DefaultResult
	h.usercase.Register(&input, &result)
	c.JSON(result.Code, result)
}

// registerProduct is the register product handler
func (h *Rest) registerProduct(c *gin.Context) {
	var input dto.RegisterProduct
	if !h.bindInput(c, &input) {
		return
	}
	var result dto.DefaultResult
	h.usercase.Register(&input, &result)
	c.JSON(result.Code, result)
}

// bindInput binds the input
func (h *Rest) bindInput(c *gin.Context, input interface{}) bool {
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 409, "message": "invalid json structure"})
		return false
	}
	return true
}
