package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/internal/client/core/dto"
	"github.com/lavinas/keel/internal/client/core/port"
	"github.com/lavinas/keel/pkg/gin_wrapper"
)

// HandlerRest is a handler for gin framework
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
	h.gin.POST("/client/insert", h.Insert)
	h.gin.POST("/client/update/:id", h.Update)
	h.gin.GET("/client/find", h.Find)
	h.gin.GET("/client/get/:param", h.Get)
}

// Run runs the gin service
func (h *HandlerRest) Run() {
	h.MapHandlers()
	h.gin.Run()
}

// Insert responds for call of creates a new client
func (h *HandlerRest) Insert(c *gin.Context) {
	var input dto.InsertInputDto
	var output dto.InsertOutputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Infof(input, "bad request: "+err.Error())
		c.JSON(http.StatusBadRequest, h.gin.H("error", "invalid json body"))
		return
	}
	if err := h.service.Insert(&input, &output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), h.gin.H("error", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, output)
}

// Find responds for call of list clients
func (h *HandlerRest) Find(c *gin.Context) {
	var input dto.FindInputDto
	var output dto.FindOutputDto
	if err := c.ShouldBindQuery(&input); err != nil {
		h.log.Infof(input, "bad request: "+err.Error())
		c.JSON(http.StatusBadRequest, h.gin.H("error", "invalid json body"))
		return
	}
	if err := h.service.Find(&input, &output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), h.gin.H("error", err.Error()))
		return
	}
	if output.Count() == 0 {
		c.JSON(http.StatusNoContent, h.gin.H("", ""))
		return
	}
	c.JSON(http.StatusOK, output)
}

// Update responds for call of updates a client
func (h *HandlerRest) Update(c *gin.Context) {
	var input dto.UpdateInputDto
	var output dto.UpdateOutputDto
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Infof(input, "bad request: "+err.Error())
		c.JSON(http.StatusBadRequest, h.gin.H("error", "invalid json body"))
		return
	}
	if err := h.service.Update(id, &input, &output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), h.gin.H("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, output)
}

// Get responds for call of get a client
func (h *HandlerRest) Get(c *gin.Context) {
	var input dto.InsertInputDto
	var output dto.InsertOutputDto
	param := c.Param("param")
	if err := h.service.Get(param, &input, &output); err != nil {
		c.JSON(h.gin.MapError(err.Error()), h.gin.H("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, output)
}
