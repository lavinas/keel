package hdlr

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/internal/client/adapter/dto"
	"github.com/lavinas/keel/internal/client/core/port"
	"github.com/lavinas/keel/pkg/gin_wrapper"
)

// HandlerGin is a handler for gin framework
type HandlerGin struct {
	log     port.Log
	service port.Service
	gin     *gin_wrapper.GinEngineWrapper
}

// NewHandlerGin creates a new HandlerGin
func NewHandlerGin(log port.Log, service port.Service) *HandlerGin {
	r := gin_wrapper.NewGinEngineWrapper(log)
	h := HandlerGin{
		log:     log,
		service: service,
		gin:     r,
	}
	return &h
}

// MapHandlers maps the handlers
func (h *HandlerGin) MapHandlers() {
	h.gin.GET("/client/list", h.ClientList)
	h.gin.POST("/client/create", h.ClientCreate)
}

// Run runs the gin service
func (h *HandlerGin) Run() {
	h.gin.Run()
	h.gin.ShutDown()
}

// ClientCreate responds for call of creates a new client
func (h *HandlerGin) ClientCreate(c *gin.Context) {
	var input dto.ClientCreateInputDto
	var output dto.ClientCreateOutputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.ClientCreate(&input, &output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), gin_wrapper.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

// ClientList responds for call of list clients
func (h *HandlerGin) ClientList(c *gin.Context) {
	var output dto.ClientListOutputDto
	if err := h.service.ClientList(&output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), gin_wrapper.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
