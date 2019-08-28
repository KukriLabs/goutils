package null_test

import (
	"encoding/json"
	"testing"

	"github.com/jmoiron/sqlx/types"
	"github.com/kukrilabs/goutils/null"
	"github.com/stretchr/testify/assert"
)

func TestNullJSONText_MarshalJSON(t *testing.T) {
	tt := []struct {
		name           string
		input          *null.NullJSONText
		expectedOutput []byte
		expectedErr    error
	}{
		{name: "Invalid NullJSONText", input: &null.NullJSONText{}, expectedOutput: []byte("null"), expectedErr: nil},
		{name: "Valid NullJSONText", input: &null.NullJSONText{
			types.NullJSONText{
				Valid:    true,
				JSONText: []byte(`{"key": "value"}`),
			},
		}, expectedOutput: []byte(`{"key": "value"}`), expectedErr: nil},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output, err := tc.input.MarshalJSON()
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedOutput, output)
			}
		})
	}
}

func TestNullJSONText_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		name           string
		input          []byte
		expectedOutput null.NullJSONText
		expectedErr    bool
	}{
		{name: "Valid JSON", input: []byte(`{"key": "value"}`), expectedOutput: null.NullJSONText{
			types.NullJSONText{
				Valid:    true,
				JSONText: []byte(`{"key": "value"}`),
			},
		}, expectedErr: false},
		{"Invalid JSON", []byte("[failure JSON]"), null.NullJSONText{}, true},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var output null.NullJSONText
			err := json.Unmarshal(tc.input, &output)
			if tc.expectedErr {
				assert.False(t, output.Valid)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedOutput, output)
			}
		})
	}
}
