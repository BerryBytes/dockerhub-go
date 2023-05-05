package dockerhub

import (
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

// String returns a pointer to a string for configuration.
func String(s string) *string {
	return &s
}

// StringValue returns the value of a String pointer
func StringValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
