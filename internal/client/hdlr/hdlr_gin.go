package hdlr

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/internal/client/core/dto"
	"github.com/lavinas/keel/internal/client/core/port"
)

// HandlerGin is a handler for gin framework
type HandlerGin struct {
	log     port.Log
	service port.Service
	gin     *gin.Engine
}

// NewHandlerGin creates a new HandlerGin
func NewHandlerGin(log port.Log, service port.Service) *HandlerGin {
	r := ginConf(log)
	h := HandlerGin{
		log:     log,
		service: service,
		gin:     r,
	}
	r.POST("/client/create", h.Create)
	return &h
}

// Run runs the gin service
func (h *HandlerGin) Run() {
	srv := ginRun(h.log, h.gin)
	ginShutDown(h.log, srv)
}

// Create responds for call of creates a new client
func (h *HandlerGin) Create(c *gin.Context) {
	var input dto.CreateInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := h.service.Create(input)
	if err != nil {
		c.JSON(mapError(err.Error()), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

// Gin assistant functions

// ginConf configures gin framework
func ginConf(l port.Log) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(l.GetFile())
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	return r
}

// ginRun runs gin service
func ginRun(l port.Log, r http.Handler) *http.Server {
	l.Info("starting gin service at 127.0.0.1:8081")
	srv := &http.Server{Addr: ":8081", Handler: r}
	quit := make(chan os.Signal, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Error("listenner error: " + err.Error())
			quit <- syscall.SIGTERM
		}
	}()
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	return srv
}

// ginShutDown shutdowns gin service
func ginShutDown(l port.Log, srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Error("server Shutdown Error: " + err.Error())
	}
	<-ctx.Done()
	l.Info("closed gin service at 127.0.0.1:8081")
}

func mapError(message string) int {
	if strings.Contains(message, "bad request") {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
