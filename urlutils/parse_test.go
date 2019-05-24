package urlutils_test

import (
	"net/url"
	"testing"

	"github.com/kukrilabs/goutils/urlutils"
	"github.com/stretchr/testify/assert"
)

func TestMustParseURL(t *testing.T) {
	tt := []struct {
		name           string
		inputURL       string
		expectedOutput *url.URL
		expectedPanic  bool
	}{
		{name: "empty string", inputURL: "", expectedOutput: &url.URL{}, expectedPanic: false},
		{"invalid scheme", ":////scheme:39:39", nil, true},
		{"valid url", "https://www.example.com/path?query", &url.URL{
			Scheme: "https", Host: "www.example.com", Path: "/path", RawQuery: "query",
		}, false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedPanic {
				assert.Panics(t, func() { urlutils.MustParseURL(tc.inputURL) })
			} else {
				assert.Equal(t, tc.expectedOutput, urlutils.MustParseURL(tc.inputURL))
			}
		})
	}
}
