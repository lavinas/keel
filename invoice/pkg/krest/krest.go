package krest

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	invoice_port = "KEEL_INVOICE_PORT"
)

// Config is the config interface for the rest handler
type Config interface {
	Get(string) string
}

// Logger is the logger interface for the rest handler
type Logger interface {
	GetFile() *os.File
	Info(string)
	Infof(format string, args ...interface{})
	Error(error)
}

// HandlerMap is the map of handlers
// First key is the http method (GET, POST, PUT, DELETE)
// Second key is the path
// Value is the handler function
type HandlerMap map[string]map[string]gin.HandlerFunc

// Rest is the rest handler for the application
type Krest struct {
	engine *gin.Engine
	config Config
	logger Logger
}

// NewRest creates a new rest handler
func NewKrest(config Config, logger Logger) *Krest {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(logger.GetFile())
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	return &Krest{
		engine: r,
		config: config,
		logger: logger,
	}
}

// Run runs the rest handler
func (g *Krest) Run(handlers HandlerMap) {
	g.logger.Info("starting rest handler")
	g.registerRoutes(handlers)
	srv := g.startServer()
	g.shutServer(srv)
	g.logger.Info("closing rest handler")
}

// registerGets registers all the get routes
func (g *Krest) registerRoutes(handlers HandlerMap) {
	for t, h := range handlers {
		for p, f := range h {
			g.logger.Infof("registering %s route %s", t, p)
			switch t {
			case "GET":
				g.engine.GET(p, f)
			case "POST":
				g.engine.POST(p, f)
			case "PUT":
				g.engine.PUT(p, f)
			case "DELETE":
				g.engine.DELETE(p, f)
			}
		}
	}
}

// startServer starts the gin server
func (g *Krest) startServer() *http.Server {
	port := g.config.Get(invoice_port)
	if port == "" {
		port = "8081"
	}
	g.logger.Infof("starting service at 127.0.0.1:%s", port)
	srv := &http.Server{Addr: ":8081", Handler: g.engine}
	quit := make(chan os.Signal, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			g.logger.Error(err)
			quit <- syscall.SIGTERM
		}
	}()
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	return srv
}

// shutServer shuts down the gin server
func (g *Krest) shutServer(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		g.logger.Error(err)
	}
	<-ctx.Done()
	g.logger.Infof("closing service at %s", srv.Addr)
}
