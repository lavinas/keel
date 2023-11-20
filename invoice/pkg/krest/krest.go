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
}

// NewRest creates a new rest handler
func NewKrest() *Krest {
	return &Krest{}
}

// Run runs the rest handler
func (g *Krest) Run(logger Logger, handlers HandlerMap) {
	logger.Info("starting rest handler")
	engine := g.getEngine(logger)
	g.registerRoutes(engine, logger, handlers)
	srv := g.startServer(engine, logger)
	g.shutServer(srv, logger)
	logger.Info("closing rest handler")
}

// getEngine sets and returns the gin engine
func (g *Krest) getEngine(logger Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	if file := logger.GetFile(); file != nil {
		gin.DefaultWriter = io.MultiWriter(logger.GetFile())
	}
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	return r
}

// registerGets registers all the get routes
func (g *Krest) registerRoutes(r *gin.Engine, logger Logger, handlers HandlerMap) {
	for t, h := range handlers {
		for p, f := range h {
			logger.Infof("registering %s route %s", t, p)
			switch t {
			case "GET":
				r.GET(p, f)
			case "POST":
				r.POST(p, f)
			case "PUT":
				r.PUT(p, f)
			case "DELETE":
				r.DELETE(p, f)
			}
		}
	}
}

// startServer starts the gin server
func (g *Krest) startServer(engine *gin.Engine, logger Logger) *http.Server {
	logger.Info("starting service at 127.0.0.1:8081")
	srv := &http.Server{Addr: ":8081", Handler: engine}
	quit := make(chan os.Signal, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(err)
			quit <- syscall.SIGTERM
		}
	}()
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	return srv
}

// shutServer shuts down the gin server
func (g *Krest) shutServer(srv *http.Server, logger Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(err)
	}
	<-ctx.Done()
	logger.Info("closing service at 127.0.0.1:8081")
}
