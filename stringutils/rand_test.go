package stringutils_test

import (
	"testing"

	"github.com/kukrilabs/goutils/stringutils"
	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	tt := []struct {
		name               string
		randomStringLength int
	}{
		{name: "Good string", randomStringLength: 5},
		{"no length", 0},
		{"mid length", 20},
		{"longer length", 100},
		{"huge length", 10000},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.randomStringLength, len(stringutils.RandAlphaNumericString(tc.randomStringLength)))
		})
	}
}
