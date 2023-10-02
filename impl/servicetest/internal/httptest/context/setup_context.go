package context

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
)

func NewTestContext(method string, endpoinURL string, request any) (echo.Context, *httptest.ResponseRecorder, error) {
	e := echo.New()
	requestJSON, err := json.Marshal(request)

	req := httptest.NewRequest(http.MethodPost, endpoinURL, bytes.NewReader(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec, err
}
