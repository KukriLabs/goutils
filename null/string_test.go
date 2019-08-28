package null_test

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/kukrilabs/goutils/null"
	"github.com/stretchr/testify/assert"
)

func TestNullString_MarshalJSON(t *testing.T) {
	tt := []struct {
		name           string
		input          *null.NullString
		expectedOutput []byte
		expectedErr    error
	}{
		{name: "Invalid NullString", input: &null.NullString{}, expectedOutput: []byte("null"), expectedErr: nil},
		{"Valid NullString", &null.NullString{
			sql.NullString{
				Valid:  true,
				String: "foobar",
			},
		}, []byte(`"foobar"`), nil},
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

func TestNullString_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		name           string
		input          []byte
		expectedOutput null.NullString
		expectedErr    bool
	}{
		{name: "Valid JSON", input: []byte(`"foobar"`), expectedOutput: null.NullString{
			sql.NullString{
				Valid:  true,
				String: "foobar",
			},
		}, expectedErr: false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var output null.NullString
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
