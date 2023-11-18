package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/invoice/internal/core/dto"
	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/gin_wrapper"
)

type HandlerRest struct {
	log     port.Log
	service port.Service
	gin     *gin_wrapper.GinEngineWrapper
}

// NewHandlerGin creates a new HandlerRest
func NewHandlerRest(log port.Log, service port.Service) *HandlerRest {
	r := gin_wrapper.NewGinEngineWrapper(log)
	h := HandlerRest{
		log:     log,
		service: service,
		gin:     r,
	}
	return &h
}

// MapHandlers maps the handlers
func (h *HandlerRest) MapHandlers() {
	h.gin.POST("/invoice/create", h.Create)
}

// Run runs the gin service
func (h *HandlerRest) Run() {
	h.MapHandlers()
	h.gin.Run()
}

// Create responds for call of creates a new invoice
func (h *HandlerRest) Create(c *gin.Context) {
	var input dto.CreateInputDto
	var output dto.CreateOutputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Infof(input, "bad request: "+err.Error())
		c.JSON(http.StatusBadRequest, h.gin.H("error", "invalid json body"))
		return
	}
	if err := h.service.Create(&input, &output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), h.gin.H("error", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, output)
}
