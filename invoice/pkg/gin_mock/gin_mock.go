package ginmock

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// GinMock is a mock of a rest gin server
type GinMock struct {
	port string
	handlerName string
	handlerType string
	output string
	srv *http.Server
	httpServerExitDone *sync.WaitGroup
}

// NewGinMock creates a new GinMock
func NewGinMock(port string, handlerName string, handlerType string, output string) *GinMock {
	return &GinMock{
		port: port,
		handlerName: handlerName,
		handlerType: handlerType,
		output: output,
		srv: &http.Server{Addr: fmt.Sprintf(":%s", port)},
		httpServerExitDone: &sync.WaitGroup{},
	}
}

// Start starts the http server
func (g *GinMock) Start() {
	g.httpServerExitDone.Add(1)
	http.HandleFunc("/", g.do)
    go func() {
        defer g.httpServerExitDone.Done()
        if err := g.srv.ListenAndServe(); err != http.ErrServerClosed {
            panic(err)
        }
    }()
}

// do is the handler function
func (g *GinMock) do(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world\n")
}

// Stop stops the http server gracefully
func (g *GinMock) Stop() {
	if err := g.srv.Shutdown(context.TODO()); err != nil {
        panic(err) // failure/timeout shutting down the server gracefully
    }
	g.httpServerExitDone.Wait()
}

