package httputils

import (
	"net/http"

	"github.com/kukrilabs/goutils/jsonutils"
)

func NotImplemented(w http.ResponseWriter) {
	jsonutils.JSONify(w, http.StatusNotImplemented, map[string]interface{}{"message": "Not implemented"})
}
