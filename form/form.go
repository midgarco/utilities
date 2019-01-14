package form

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// Form object attempts to parse all the potential inputs
// that a request can use to pass data
type Form struct {
	data map[string]interface{}
}

const (
	maxUpload = 32 // in MB
)

// Parse ...
func Parse(r *http.Request) *Form {
	f := &Form{data: make(map[string]interface{})}

	// post data
	fm := ParseForm(r)
	for k, vs := range fm.data {
		f.data[k] = vs
	}

	// json
	jd := ParseJSON(r)
	for k, vs := range jd.data {
		f.data[k] = vs
	}

	// query string
	qs := ParseQueryString(r)
	for k, vs := range qs.data {
		f.data[k] = vs
	}

	return f
}

// ParseForm parses the posted data
func ParseForm(r *http.Request) *Form {
	f := &Form{data: make(map[string]interface{})}
	r.ParseMultipartForm(maxUpload << 20)
	for k, vs := range r.Form {
		f.data[k] = strings.Join(vs, ",")
	}
	return f
}

// ParseQueryString parses the query string parameters
func ParseQueryString(r *http.Request) *Form {
	f := &Form{data: make(map[string]interface{})}
	qs, _ := url.ParseQuery(r.URL.RawQuery)
	for k, vs := range qs {
		f.data[k] = strings.Join(vs, ",")
	}
	return f
}

// ParseJSON parses the request body for the parameters
func ParseJSON(r *http.Request) *Form {
	f := &Form{data: make(map[string]interface{})}
	jsonData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&jsonData)
	for k, vs := range jsonData {
		f.data[k] = vs
	}
	return f
}

// Get the key from the form data
func (f Form) Get(k string) string {
	if _, ok := f.data[k]; ok {
		return f.data[k].(string)
	}
	return ""
}

// GetSlice returns the key as a slice of strings
func (f Form) GetSlice(k string) []string {
	s := reflect.ValueOf(f.data[k])
	if s.Kind() != reflect.Slice {
		return nil
	}
	ret := make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface().(string)
	}
	return ret
}
