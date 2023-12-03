package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/invoice/internal/core/domain"
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
		c.JSON(http.StatusBadRequest, gin.H{"code": 409, "message": "invalid json structure"})
		return
	}
	var result dto.DefaultResult
	h.usercase.Create(obj, &result)
	c.JSON(result.Code, result)
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
