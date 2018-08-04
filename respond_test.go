package httputils

import (
	// "log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type teststruct struct {
	Name string `json:"name"`
}

func TestDecodeBody(t *testing.T) {
	t.Parallel()
	json := `{"name":"creativeName"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	var r teststruct
	DecodeBody(req, &r)
	assert.Equal(t, "creativeName", r.Name)
}

func TestEncodeBody(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	r := teststruct{
		Name: "superCreative",
	}
	EncodeBody(rec, req, &r)
	assert.JSONEq(t, `{"name":"superCreative"}`, rec.Body.String())
}

func TestRespond(t *testing.T) {
	t.Parallel()
	t.Skip("Test not implemented")
}

func TestRespondErr(t *testing.T) {
	t.Parallel()
	t.Skip("Test not implemented")
}

func TestRespondHTTPErr(t *testing.T) {
	t.Parallel()
	t.Skip("Test not implemented")
}
