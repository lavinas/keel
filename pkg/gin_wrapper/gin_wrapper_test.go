package gin_wrapper

import (
	"net/http"
	"testing"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func TestRun(t *testing.T) {
	log := LogMock{}
	g := NewGinEngineWrapper(&log)
	go g.Run()
	time.Sleep(time.Second * 1)
	if log.mtype[0] != "Info" {
		t.Errorf("Error: %s", log.mtype)
	}
	if !strings.Contains(log.msgs[0], "starting gin service at") {
		t.Errorf("Error: %s", log.msgs[0])
	}
}

func TestRunListenError(t *testing.T) {
	log := LogMock{}
	g := NewGinEngineWrapper(&log)
	go g.Run()
	time.Sleep(time.Second * 1)
	log2 := LogMock{}
	g2 := NewGinEngineWrapper(&log2)
	g2.Run()
	time.Sleep(time.Second * 1)
	if log2.mtype[1] != "Error" {
		t.Errorf("Error: %s", log2.mtype)
	}
	if !strings.Contains(log2.msgs[1], "listenner error:") {
		t.Errorf("Error: %s", log2.msgs[0])
	}
}

func TestPOST(t *testing.T) {
	log := LogMock{}
	g := NewGinEngineWrapper(&log)
	go g.Run()
	time.Sleep(time.Second * 1)
	g.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"test": "test"})
	})
}

func TestPUST(t *testing.T) {
	log := LogMock{}
	g := NewGinEngineWrapper(&log)
	go g.Run()
	time.Sleep(time.Second * 1)
	g.PUT("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"test": "test"})
	})
}

func TestGET(t *testing.T) {
	log := LogMock{}
	g := NewGinEngineWrapper(&log)
	go g.Run()
	time.Sleep(time.Second * 1)
	g.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"test": "test"})
	})
}

func TestDELETE(t *testing.T) {
	log := LogMock{}
	g := NewGinEngineWrapper(&log)
	go g.Run()
	time.Sleep(time.Second * 1)
	g.DELETE("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"test": "test"})
	})
}


func TestMapError(t *testing.T) {
	g := NewGinEngineWrapper(&LogMock{})
	if g.MapError("bad request: invalid x") != http.StatusBadRequest {
		t.Errorf("Error: %s", "error")
	}
	if g.MapError("not found: invalid x") != http.StatusNotFound {
		t.Errorf("Error: %s", "error")
	}
	if g.MapError("conflict: invalid x") != http.StatusConflict {
		t.Errorf("Error: %s", "error")
	}
	if g.MapError("unauthorized: invalid x") != http.StatusUnauthorized {
		t.Errorf("Error: %s", "error")
	}
	if g.MapError("invalid x") != http.StatusInternalServerError {
		t.Errorf("Error: %s", "error")
	}
}

func TestH(t *testing.T) {
	g := NewGinEngineWrapper(&LogMock{})
	if g.H("test", "test")["test"] != "test" {
		t.Errorf("Error: %s", "error")
	}
}

