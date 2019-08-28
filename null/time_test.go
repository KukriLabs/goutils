package null_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/kukrilabs/goutils/null"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestNullPQTime_MarshalJSON(t *testing.T) {
	tt := []struct {
		name           string
		input          *null.NullPQTime
		expectedOutput []byte
		expectedErr    error
	}{
		{name: "Invalid NullTime", input: &null.NullPQTime{}, expectedOutput: []byte("null"), expectedErr: nil},
		{"Valid NullTime", &null.NullPQTime{
			pq.NullTime{
				Valid: true,
				Time:  time.Date(2001, time.January, 1, 2, 0, 0, 0, time.UTC),
			},
		}, []byte(`"2001-01-01T02:00:00Z"`), nil},
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

func TestNullPQTime_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		name           string
		input          []byte
		expectedOutput null.NullPQTime
		expectedErr    bool
	}{
		{name: "Valid JSON", input: []byte(`"2001-01-01T02:00:00Z"`), expectedOutput: null.NullPQTime{
			pq.NullTime{
				Valid: true,
				Time:  time.Date(2001, time.January, 1, 2, 0, 0, 0, time.UTC),
			},
		}, expectedErr: false},
		{"Invalid JSON", []byte(`i'm a failure`), null.NullPQTime{}, true},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var output null.NullPQTime
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
