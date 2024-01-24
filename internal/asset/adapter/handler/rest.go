package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/internal/asset/core/dto"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
	"github.com/lavinas/keel/pkg/krest"
)

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
			"/asset/ping": h.ping,
		},
		"POST": {
			"/asset/tax":       h.Create,
			"/asset/class":     h.Create,
			"/asset/asset":     h.Create,
			"/asset/portfolio": h.Create,
			"/asset/statement": h.Create,
		},
	}
	h.krest.Run(handlers)
}

// ping is the ping handler
func (h *Rest) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Create is the create handler
func (h *Rest) Create(c *gin.Context) {
	in, out := h.createFactory(c)
	if in == nil || out == nil {
		err := kerror.NewKError(kerror.BadRequest, "invalid url")
		h.logger.Error(err)
		rError := NewError(err)
		c.JSON(rError.Status, rError)
		return
	}
	if err := c.ShouldBindJSON(in); err != nil {
		h.logger.Error(err)
		rError := NewError(kerror.NewKError(kerror.BadRequest, "invalid json structure"))
		c.JSON(rError.Status, rError)
		return
	}
	if err := h.usercase.Create(in, out); err != nil {
		h.logger.Error(err)
		rError := NewError(err)
		c.JSON(rError.Status, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

// createFactory creates a new dto for input and output
func (h *Rest) createFactory(c *gin.Context) (port.CreateDtoIn, port.CreateDtoOut) {
	switch c.Request.URL.Path {
	case "/asset/tax":
		return &dto.TaxCreateIn{}, &dto.TaxCreateOut{}
	case "/asset/class":
		return &dto.ClassCreateIn{}, &dto.ClassCreateOut{}
	case "/asset/asset":
		return &dto.AssetCreateIn{}, &dto.AssetCreateOut{}
	case "/asset/portfolio":
		return &dto.PortfolioCreateIn{}, &dto.PortfolioCreateOut{}
	case "/asset/statement":
		return &dto.StatementCreateIn{}, &dto.StatementCreateOut{}
	default:
		return nil, nil
	}
}
