package gin_wrapper

import (
	"testing"
	"net/http"
)

func TestMapError(t *testing.T) {
	g := GinEngineWrapper{}
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