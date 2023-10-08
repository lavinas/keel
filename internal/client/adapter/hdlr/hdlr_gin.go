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
	h.gin.POST("/client/insert", h.ClientInsert)
	h.gin.POST("/client/update/:id", h.ClientUpdate)
	h.gin.GET("/client/get/:param", h.ClientGet)
}

// Run runs the gin service
func (h *HandlerGin) Run() {
	h.gin.Run()
	h.gin.ShutDown()
}

// ClientInsert responds for call of creates a new client
func (h *HandlerGin) ClientInsert(c *gin.Context) {
	var input dto.ClientInsertInputDto
	var output dto.ClientInserOutputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Infof(input, "bad request: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}
	if err := h.service.ClientInsert(&input, &output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), gin_wrapper.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

// ClientList responds for call of list clients
func (h *HandlerGin) ClientList(c *gin.Context) {
	var input dto.ClientListInputDto
	var output dto.ClientListOutputDto
	if err := c.ShouldBindQuery(&input); err != nil {
		h.log.Infof(input, "bad request: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}
	if err := h.service.ClientList(&input, &output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), gin_wrapper.H{"error": err.Error()})
		return
	}
	if output.Count() == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, output)
}

// ClientUpdate responds for call of updates a client
func (h *HandlerGin) ClientUpdate(c *gin.Context) {
	var input dto.ClientInsertInputDto
	var output dto.ClientInserOutputDto
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Infof(input, "bad request: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}
	if err := h.service.ClientUpdate(id, &input, &output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), gin_wrapper.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

// ClientGet responds for call of get a client
func (h *HandlerGin) ClientGet(c *gin.Context) {
	var input dto.ClientInsertInputDto
	var output dto.ClientInserOutputDto
	param := c.Param("param")
	if err := h.service.ClientGet(param, &input, &output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), gin_wrapper.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
