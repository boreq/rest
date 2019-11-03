package rest

import "net/http"

type Response struct {
	header     http.Header
	statusCode int
	body       interface{}
}

func NewResponse(body interface{}) Response {
	return Response{
		header:     make(http.Header),
		statusCode: 200,
		body:       body,
	}
}

func (r Response) Header() http.Header {
	return r.clone().header
}

func (r Response) StatusCode() int {
	return r.clone().statusCode
}

func (r Response) Body() interface{} {
	return r.clone().body
}

func (r Response) WithStatusCode(statusCode int) Response {
	r2 := r.clone()
	r2.statusCode = statusCode
	return r2
}

func (r Response) WithHeader(key, value string) Response {
	r2 := r.clone()
	r2.header.Add(key, value)
	return r2
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
