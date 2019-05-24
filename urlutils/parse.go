package urlutils

import "net/url"

func MustParseURL(u string) *url.URL {
	output, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	return output
}
