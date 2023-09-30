package hdlr

import (
	"os"
	"io"
	"bytes"
	"net/http"
	"net/url"
	"net/http/httptest"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/internal/client/core/port"
)

//mock gin context
func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
	Header: make(http.Header),
	URL:    &url.URL{},
	}

	return ctx
}

//mock postrequest
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
func MockJsonGet(c *gin.Context) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)

	// set query params
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
}

func (s *ServiceMock) Create(input port.CreateInputDto, output port.CreateOutputDto) error {
	output.Fill("1", "name", "nickname", "document", "phone", "email")
	return nil
}
