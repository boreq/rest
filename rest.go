// Package rest implements a framework for creating a JSON REST API. It
// automatically serializes the responses as JSON  and provides the user with
// the ability to implement a clearer flow of control within the handler thanks
// to encapsulating all calls to the response writer within a single return
// value from the handler.
//
//  func handler(r *http.Request) rest.RestResponse {
//  	value, err := someFunction()
//  	if err != nil {
//  		return rest.InternalServerError
//  	}
//  	return rest.NewResponse(value)
//  }
package rest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/boreq/errors"
)

// RestResponse encapsulates parameters passed to the response writer.
type RestResponse interface {
	Header() http.Header
	StatusCode() int
	Body() interface{}
}

// HandlerFunc returns a rest response instead of accepting a response writer
// to simplify the flow of control within the handler and replace multiple
// calls to the response writer with a single return statement.
type HandlerFunc func(r *http.Request) RestResponse

// Call executes the provided handler passing it the provided request. The
// response returned from the handler is written to the provided response
// writer. In most cases you should be using Wrap instead of calling this
// method directly.
func Call(w http.ResponseWriter, r *http.Request, handler HandlerFunc) error {
	response := handler(r)

	code := response.StatusCode()
	j, err := json.Marshal(response.Body())
	if err != nil {
		j, err = json.Marshal(ErrInternalServerError.Body())
		if err != nil {
			return errors.Wrap(err, "could not marshal the builtin error type")
		}
		code = ErrInternalServerError.StatusCode()
	}

	// write the headers
	for key, values := range response.Header() {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
	}

	// write the status and the body
	w.WriteHeader(code)
	_, err = bytes.NewBuffer(j).WriteTo(w)
	return errors.Wrap(err, "writing the response failed")
}

// Wrap encapsulates the provided handler to make it compatibile with net/http.
func Wrap(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Call(w, r, handler); err != nil {
			log.Printf("the rest library encountered an error and was unable to handle it: %s", err)
		}
	}
}
