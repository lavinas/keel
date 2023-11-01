package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/client/internal/core/port"
)

// mock gin context
func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	return ctx
}

// mock postrequest
func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

// mock GET request
func MockJsonGet(c *gin.Context, content map[string]string) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	for k, v := range content {
		c.Params = append(c.Params, gin.Param{Key: k, Value: v})
	}
	u := url.Values{}
	u.Add("skip", "5")
	u.Add("limit", "10")
	c.Request.URL.RawQuery = u.Encode()
}

// mock PUT request
func MockJsonPut(c *gin.Context, content interface{}, params gin.Params) {
	c.Request.Method = "PUT"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)
	c.Params = params

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

// mock DELETE request
func MockJsonDelete(c *gin.Context, params gin.Params) {
	c.Request.Method = "DELETE"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)
	c.Params = params
}

// Log Mock
type LogMock struct {
	mtype string
	msg   string
}

func (l *LogMock) GetFile() *os.File {
	return nil
}
func (l *LogMock) Info(msg string) {
	l.mtype = "Info"
	l.msg = msg
}
func (l *LogMock) Infof(input any, message string) {
	l.mtype = "Info"
	b, _ := json.Marshal(input)
	l.Info(message + " | " + string(b))
}
func (l *LogMock) Error(msg string) {
	l.mtype = "Error"
	l.msg = msg
}
func (l *LogMock) Errorf(input any, err error) {
	b, _ := json.Marshal(input)
	l.Error(err.Error() + " | " + string(b))
}
func (l *LogMock) Close() {
}

// Service Mock
type ServiceMock struct {
	Status string
}

func (s *ServiceMock) Insert(input port.InsertInputDto, output port.InsertOutputDto) error {
	if s.Status == "ok" {
		name, nick, doc, phone, email := input.Get()
		output.Fill("1", name, nick, doc, phone, email)
		return nil
	}
	if s.Status == "bad request" {
		return errors.New("bad request: invalid json body")
	}
	if s.Status == "internal error" {
		return errors.New("internal error")
	}
	return nil
}
func (s *ServiceMock) Find(input port.FindInputDto, output port.FindOutputDto) error {
	if s.Status == "no content" {
		return nil
	}
	if s.Status == "ok" {
		output.Append("1", "Jose da Silva", "jose_da_silva_222", "206.656.600-49", "+5511999999999", "test@test.com")
		output.SetPage(1, 10)
		return nil
	}
	if s.Status == "bad request" {
		return errors.New("bad request: invalid xxxx")
	}
	return nil
}
func (s *ServiceMock) Update(id string, input port.UpdateInputDto, output port.UpdateOutputDto) error {
	if s.Status == "ok" {
		name, nick, doc, phone, email := input.Get()
		output.Fill(id, name, nick, doc, phone, email)
		return nil
	}
	if s.Status == "bad request" {
		return errors.New("bad request: invalid json body")
	}
	return nil
}
func (s *ServiceMock) Get(param string, input port.InsertInputDto, output port.InsertOutputDto) error {
	if s.Status == "ok" {
		name, nick, doc, phone, email := input.Get()
		output.Fill("1", name, nick, doc, phone, email)
		return nil
	}
	if s.Status == "bad request" {
		return errors.New("bad request: invalid json body")
	}
	return nil
}

// GinEngineWrapper Mock
type GinEngineWrapperMock struct {
	Maps map[string]interface{}
}

func NewGinEngineWrapperMock() *GinEngineWrapperMock {
	return &GinEngineWrapperMock{
		Maps: make(map[string]interface{}),
	}
}

func (g *GinEngineWrapperMock) Run() *http.Server {
	return nil
}
func (g *GinEngineWrapperMock) ShutDown()               {}
func (g *GinEngineWrapperMock) MapError(message string) {}
func (g *GinEngineWrapperMock) POST(relativePath string, handlers ...gin.HandlerFunc) {
	g.Maps["relativePath"] = handlers
}
func (g *GinEngineWrapperMock) PUT(relativePath string, handlers ...gin.HandlerFunc) {
	g.Maps["relativePath"] = handlers
}
func (g *GinEngineWrapperMock) GET(relativePath string, handlers ...gin.HandlerFunc) {
	g.Maps["relativePath"] = handlers
}
func (g *GinEngineWrapperMock) DELETE(relativePath string, handlers ...gin.HandlerFunc) {
	g.Maps["relativePath"] = handlers
}
