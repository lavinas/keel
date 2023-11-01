package gin_wrapper

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
	"github.com/lavinas/keel/util/pkg/log"
)

// GinEngineWrapper is a wrapper handler for gin framework with graceful configuration and shutdown
type GinEngineWrapper struct {
	log    *log.Log
	engine *gin.Engine
	server *http.Server
}

// NewGinWrapper creates a new GinEngine
func NewGinEngineWrapper(log *log.Log) *GinEngineWrapper {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(log.GetFile())
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	return &GinEngineWrapper{
		log:    log,
		engine: r,
	}
}

// Run runs the gin service
func (g *GinEngineWrapper) Run() {
	g.log.Info("starting gin service at 127.0.0.1:8081")
	srv := &http.Server{Addr: ":8081", Handler: g.engine}
	quit := make(chan os.Signal, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			g.log.Error("listenner error: " + err.Error())
			quit <- syscall.SIGTERM
		}
	}()
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	g.server = srv
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := g.server.Shutdown(ctx); err != nil {
		g.log.Error("server Shutdown Error: " + err.Error())
	}
	<-ctx.Done()
	g.log.Info("closing gin service at 127.0.0.1:8081")
}

// MapError maps error message to http status code
func (g *GinEngineWrapper) MapError(message string) int {
	var mmap = map[string]int{
		"not found":    http.StatusNotFound,
		"bad request":  http.StatusBadRequest,
		"conflict":     http.StatusConflict,
		"unauthorized": http.StatusUnauthorized,
		"no content":   http.StatusNoContent,
	}
	m := strings.Split(message, ":")[0]
	if v, ok := mmap[m]; ok {
		return v
	}
	return http.StatusInternalServerError
}

// POST is a wrapper for gin POST
func (g *GinEngineWrapper) POST(relativePath string, handlers ...gin.HandlerFunc) {
	g.engine.POST(relativePath, handlers...)
}

// PUT is a wrapper for gin PUT
func (g *GinEngineWrapper) PUT(relativePath string, handlers ...gin.HandlerFunc) {
	g.engine.PUT(relativePath, handlers...)
}

// GET is a wrapper for gin PUT
func (g *GinEngineWrapper) GET(relativePath string, handlers ...gin.HandlerFunc) {
	g.engine.GET(relativePath, handlers...)
}

// DELETE is a wrapper for gin PUT
func (g *GinEngineWrapper) DELETE(relativePath string, handlers ...gin.HandlerFunc) {
	g.engine.DELETE(relativePath, handlers...)
}

// H is a shortcut for map[string]interface{} as gin
func (g *GinEngineWrapper) H(idx string, ctx any) map[string]any {
	return map[string]any{
		idx: ctx,
	}
}
