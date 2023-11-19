package handler

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Rest is the rest handler for the application
type Rest struct {
	engine *gin.Engine
}

// NewRest creates a new rest handler
func NewRest(logger port.Logger) *Rest {
	engine := GetEngine(logger)
	return &Rest{
		engine: engine,
	}
}

func GetEngine(logger port.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(logger.GetFile())
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	return r
}
