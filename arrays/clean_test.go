package arrays_test

import (
	"testing"

	"github.com/kukrilabs/goutils/arrays"
	"github.com/stretchr/testify/assert"
)

func TestTrimEmptyStrings(t *testing.T) {
	tt := []struct {
		name     string
		input    []string
		expected []string
	}{
		{name: "empty array", input: []string{}, expected: []string{}},
		{"array of empty strings", []string{"", "  ", "\t", "\n \t \r", "    "}, []string{}},
		{"functional array", []string{"foo", "bar"}, []string{"foo", "bar"}},
		{"mixed array", []string{"foo", "    ", "bar"}, []string{"foo", "bar"}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, arrays.TrimEmptyStrings(tc.input))
		})
	}
}
