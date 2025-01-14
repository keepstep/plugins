package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type testcase struct {
	Token  string
	ReqFn  func(opts *Options) *Request
	Method string
	URI    string
	Body   interface{}
	Header map[string]string
	Assert func(req *http.Request) bool
}

type assertFn func(req *http.Request) bool

var tests = []testcase{
	{
		ReqFn: func(opts *Options) *Request {
			return NewRequest(opts).Get().Resource("services")
		},
		Method: "GET",
		URI:    "/api/v1/namespaces/default/services/",
	},
	{
		ReqFn: func(opts *Options) *Request {
			return NewRequest(opts).Get().Resource("services").Name("foo")
		},
		Method: "GET",
		URI:    "/api/v1/namespaces/default/services/foo",
	},
	{
		ReqFn: func(opts *Options) *Request {
			return NewRequest(opts).Get().Resource("services").Namespace("test").Name("bar")
		},
		Method: "GET",
		URI:    "/api/v1/namespaces/test/services/bar",
	},
	{
		ReqFn: func(opts *Options) *Request {
			return NewRequest(opts).Get().Resource("pods").Params(&Params{LabelSelector: map[string]string{"foo": "bar"}})
		},
		Method: "GET",
		URI:    "/api/v1/namespaces/default/pods/?labelSelector=foo%3Dbar",
	},
	{
		ReqFn: func(opts *Options) *Request {
			return NewRequest(opts).Post().Resource("services").Name("foo").Body(map[string]string{"foo": "bar"})
		},
		Method: "POST",
		URI:    "/api/v1/namespaces/default/services/foo",
		Body:   map[string]string{"foo": "bar"},
	},
	{
		ReqFn: func(opts *Options) *Request {
			return NewRequest(opts).Put().Resource("endpoints").Name("baz").Body(map[string]string{"bam": "bar"})
		},
		Method: "PUT",
		URI:    "/api/v1/namespaces/default/endpoints/baz",
		Body:   map[string]string{"bam": "bar"},
	},
	{
		ReqFn: func(opts *Options) *Request {
			return NewRequest(opts).Patch().Resource("endpoints").Name("baz").Body(map[string]string{"bam": "bar"})
		},
		Method: "PATCH",
		URI:    "/api/v1/namespaces/default/endpoints/baz",
		Body:   map[string]string{"bam": "bar"},
	},
	{
		ReqFn: func(opts *Options) *Request {
			return NewRequest(opts).Patch().Resource("endpoints").Name("baz").SetHeader("foo", "bar")
		},
		Method: "PATCH",
		URI:    "/api/v1/namespaces/default/endpoints/baz",
		Header: map[string]string{"foo": "bar"},
	},
}

var wrappedHandler = func(test *testcase, t *testing.T) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if len(test.Token) > 0 && (len(auth) == 0 || auth != "Bearer "+test.Token) {
			t.Errorf("test case token (%s) did not match expected token (%s)", "Bearer "+test.Token, auth)
		}

		if len(test.Method) > 0 && test.Method != r.Method {
			t.Errorf("test case Method (%s) did not match expected Method (%s)", test.Method, r.Method)
		}

		if len(test.URI) > 0 && test.URI != r.URL.RequestURI() {
			t.Errorf("test case URI (%s) did not match expected URI (%s)", test.URI, r.URL.RequestURI())
		}

		if test.Body != nil {
			var res map[string]string
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&res); err != nil {
				t.Errorf("decoding body failed: %v", err)
			}
			if !reflect.DeepEqual(res, test.Body) {
				t.Error("body did not match")
			}
		}

		if test.Header != nil {
			for k, v := range test.Header {
				if r.Header.Get(k) != v {
					t.Error("header did not exist")
				}
			}
		}

		w.WriteHeader(http.StatusOK)
	})
}

func TestRequest(t *testing.T) {
	for _, test := range tests {
		ts := httptest.NewServer(wrappedHandler(&test, t))
		req := test.ReqFn(&Options{
			Host:        ts.URL,
			Client:      &http.Client{},
			BearerToken: &test.Token,
			Namespace:   "default",
		})
		res := req.Do()
		if res.Error() != nil {
			t.Errorf("Did not expect to fail with %v", res.Error())
		}

		ts.Close()
	}
}
