package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/kerror"
	"github.com/lavinas/keel/invoice/pkg/krest"
)

const (
	CreateDetail = "created"
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
			"/invoice/client":      h.Create,
			"/invoice/instruction": h.Create,
			"/invoice/product":     h.Create,
			"invoice":              h.Create,
		},
	}
	h.krest.Run(handlers)
}

// ping is the ping handler
func (h *Rest) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Register is the register handler
func (h *Rest) Create(c *gin.Context) {
	obj := h.domainFactory(c)
	if err := c.ShouldBindJSON(obj); err != nil {
		h.logger.Error(err)
		rError := NewError(kerror.NewKError(kerror.BadRequest, "invalid json structure"))
		c.JSON(rError.Status, rError)
		return
	}
	if err := h.usercase.Create(obj); err != nil {
		rError := NewError(err)
		c.JSON(rError.Status, rError)
		return
	}
	result := NewRestResult(http.StatusCreated, CreateDetail, &obj)
	c.JSON(http.StatusOK, result)
}

// domainFactory creates a new domain object
func (h *Rest) domainFactory(c *gin.Context) port.Domain {
	domainMap := map[string]port.Domain{
		"/invoice/product":     &domain.Product{},
		"/invoice/instruction": &domain.Instruction{},
		"/invoice/client":      &domain.Client{},
		"/invoice":             &domain.Invoice{},
	}
	return domainMap[c.Request.URL.Path]
}
