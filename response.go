package rest

import "net/http"

// Response represents a REST API response and implements the RestResponse
// interface. This is the main type you should be returning from your handlers.
// Use the Error type instead of using the response directly to return API
// errors from your handler.
type Response struct {
	header     http.Header
	statusCode int
	body       interface{}
}

// NewResponse creates a new response with the provided body and status code
// http.StatusOK. Body should be a value which can be serialized to JSON, for
// example a struct.
func NewResponse(body interface{}) Response {
	return Response{
		header:     make(http.Header),
		statusCode: 200,
		body:       body,
	}
}

// WithStatusCode returns a new response with the changed status code.
func (r Response) WithStatusCode(statusCode int) Response {
	r2 := r.clone()
	r2.statusCode = statusCode
	return r2
}

// WithStatusCode returns a new response with the added response header. Adding
// the same header multiple times appends the new value instead of replacing it
// leading to multiple headers with the same name.
func (r Response) WithHeader(key, value string) Response {
	r2 := r.clone()
	r2.header.Add(key, value)
	return r2
}

// Header is a method implementing RestResponse.
func (r Response) Header() http.Header {
	return r.clone().header
}

// StatusCode is a method implementing RestResponse.
func (r Response) StatusCode() int {
	return r.clone().statusCode
}

// Body is a method implementing RestResponse.
func (r Response) Body() interface{} {
	return r.clone().body
}

func (r Response) clone() Response {
	r2 := Response{
		header:     make(http.Header),
		statusCode: r.statusCode,
		body:       r.body,
	}

	if r.header != nil {
		r2.header = r.header.Clone()
	}

	return r2
}
