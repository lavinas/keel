package hdlr

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/internal/client/core/port"
	"github.com/lavinas/keel/internal/client/core/domain"
)

type HandlerGin struct {
	log port.Log
	service port.ClientService
	gin *gin.Engine
}

func NewHandlerGin(log port.Log, service port.ClientService) *HandlerGin {
	r := ginConf(log)
	h := HandlerGin{
		log: log,
		service: service,
		gin: r,
	}
	r.POST("/create", h.Create)
	r.GET("/list", h.ListAll)
	return &h
}

func (h *HandlerGin) Run() {
	srv := ginRun(h.log, h.gin)
	ginShutDown(h.log, srv)
}

func (h *HandlerGin) Create(c *gin.Context) {
	var input domain.CreateInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := h.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

func (h *HandlerGin) ListAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}


func ginConf(l port.Log) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(l.GetFile())
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	return r
}

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

func ginShutDown(l port.Log, srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Error("server Shutdown Error: " + err.Error())
	}
	<-ctx.Done()
	l.Info("closed gin service at 127.0.0.1:8081")
}
