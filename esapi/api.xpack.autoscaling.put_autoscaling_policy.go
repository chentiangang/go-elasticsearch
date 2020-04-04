// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 8.0.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newAutoscalingPutAutoscalingPolicyFunc(t Transport) AutoscalingPutAutoscalingPolicy {
	return func(name string, body io.Reader, o ...func(*AutoscalingPutAutoscalingPolicyRequest)) (*Response, error) {
		var r = AutoscalingPutAutoscalingPolicyRequest{Name: name, Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// AutoscalingPutAutoscalingPolicy - Creates a new autoscaling policy.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-put-autoscaling-policy.html.
//
type AutoscalingPutAutoscalingPolicy func(name string, body io.Reader, o ...func(*AutoscalingPutAutoscalingPolicyRequest)) (*Response, error)

// AutoscalingPutAutoscalingPolicyRequest configures the Autoscaling Put Autoscaling Policy API request.
//
type AutoscalingPutAutoscalingPolicyRequest struct {
	Body io.Reader

	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r AutoscalingPutAutoscalingPolicyRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_autoscaling") + 1 + len("policy") + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString("_autoscaling")
	path.WriteString("/")
	path.WriteString("policy")
	path.WriteString("/")
	path.WriteString(r.Name)

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f AutoscalingPutAutoscalingPolicy) WithContext(v context.Context) func(*AutoscalingPutAutoscalingPolicyRequest) {
	return func(r *AutoscalingPutAutoscalingPolicyRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f AutoscalingPutAutoscalingPolicy) WithPretty() func(*AutoscalingPutAutoscalingPolicyRequest) {
	return func(r *AutoscalingPutAutoscalingPolicyRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f AutoscalingPutAutoscalingPolicy) WithHuman() func(*AutoscalingPutAutoscalingPolicyRequest) {
	return func(r *AutoscalingPutAutoscalingPolicyRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f AutoscalingPutAutoscalingPolicy) WithErrorTrace() func(*AutoscalingPutAutoscalingPolicyRequest) {
	return func(r *AutoscalingPutAutoscalingPolicyRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f AutoscalingPutAutoscalingPolicy) WithFilterPath(v ...string) func(*AutoscalingPutAutoscalingPolicyRequest) {
	return func(r *AutoscalingPutAutoscalingPolicyRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f AutoscalingPutAutoscalingPolicy) WithHeader(h map[string]string) func(*AutoscalingPutAutoscalingPolicyRequest) {
	return func(r *AutoscalingPutAutoscalingPolicyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f AutoscalingPutAutoscalingPolicy) WithOpaqueID(s string) func(*AutoscalingPutAutoscalingPolicyRequest) {
	return func(r *AutoscalingPutAutoscalingPolicyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}