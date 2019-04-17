package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func Test_server_SuperErrorHandler(t *testing.T) {
	s := &server{
		echo:         echo.New(),
		serverConfig: nil,
	}
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	rec.Code = 500
	s.QilinStoreErrorHandler(errors.New("Internal error"), c)

}

func TestSuperErrorHandler(t *testing.T) {
	type args struct {
		err     error
		c       echo.Context
		isDebug bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QilinStoreErrorHandler(tt.args.err, tt.args.c, tt.args.isDebug)
		})
	}
}
