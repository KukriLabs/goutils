package jsonutils_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/kukrilabs/goutils/jsonutils"
	"github.com/stretchr/testify/assert"
)

func mustMarshal(input interface{}) []byte {
	output, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	return output
}

func TestJSONify(t *testing.T) {
	tt := []struct {
		name          string
		inputResponse interface{}
		inputStatus   int
		expected      []byte
	}{
		{name: "nil response", inputResponse: nil, inputStatus: 200, expected: nil},
		{"good response", map[string]interface{}{"foo": "bar"}, 200, mustMarshal(map[string]interface{}{"foo": "bar"})},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			jsonutils.JSONify(writer, tc.inputStatus, tc.inputResponse)
			assert.Equal(t, tc.expected, bytes.TrimSpace(writer.Body.Bytes()))
			assert.Equal(t, tc.inputStatus, writer.Code)
			assert.Equal(t, "application/json", writer.Header().Get("Content-Type"))
		})
	}
}
