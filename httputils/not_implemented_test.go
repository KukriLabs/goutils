package httputils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kukrilabs/goutils/httputils"
	"github.com/stretchr/testify/assert"
)

func TestNotImplemented(t *testing.T) {
	tt := []struct {
		name                 string
		expectedResponseCode int
	}{
		{name: "everything okay", expectedResponseCode: http.StatusNotImplemented},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			httputils.NotImplemented(rw)
			assert.Equal(t, tc.expectedResponseCode, rw.Code)
		})
	}
}
