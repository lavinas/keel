package ginmock

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"strings"

	"github.com/gin-gonic/gin"
)

// GinMock is a mock of a rest gin server
type GinMock struct {
	port int
	handlerName string
	handlerType string
	statusCode int
	output interface{}
	gin *gin.Engine
	srv *http.Server
	httpServerExitDone *sync.WaitGroup
}

// NewGinMock creates a new GinMock
func NewGinMock(port int) *GinMock {
	if port == 0 {
		port = 8080
	}
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.SetTrustedProxies([]string{"127.0.0.1"})
	return &GinMock{
		port: port,
		handlerName: "/",
		handlerType: "get",
		output: map[string]string{"message": "hello world"},
		gin: g,
		srv: &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: g},
		httpServerExitDone: &sync.WaitGroup{},
	}
}

// Start starts the http server
func (g *GinMock) Start(handlerName string, handlerType string, statusCode int, output interface{}) {
	g.handlerName = handlerName
	g.handlerType = strings.ToLower(handlerType)
	g.statusCode = statusCode
	g.output = output
	g.httpServerExitDone.Add(1)
	switch g.handlerType {
		case "get": g.gin.GET(g.handlerName, g.do)
		case "post": g.gin.POST(g.handlerName, g.do)
		case "put": g.gin.PUT(g.handlerName, g.do)
		case "delete": g.gin.DELETE(g.handlerName, g.do)
		default: g.gin.GET(g.handlerName, g.do)
	}
    go func() {
        defer g.httpServerExitDone.Done()
		g.srv.ListenAndServe()
    }()
}

// do is the handler function
func (g *GinMock) do(c *gin.Context) {
	c.JSON(g.statusCode, g.output)
}

// Stop stops the http server gracefully
func (g *GinMock) Stop() {
	g.srv.Shutdown(context.TODO())
	g.httpServerExitDone.Wait()
}

