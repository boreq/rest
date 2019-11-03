// Package rest implements a framework for creating a JSON REST API.
package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/boreq/errors"
)

type RestResponse interface {
	Header() http.Header
	StatusCode() int
	Body() interface{}
}

type HandlerFunc func(r *http.Request) RestResponse

func call(w http.ResponseWriter, r *http.Request, handler HandlerFunc) error {
	response := handler(r)

	code := response.StatusCode()
	j, err := json.Marshal(response.Body())
	if err != nil {
		println("ERROR")
		j, err = json.Marshal(ErrInternalServerError.Body())
		if err != nil {
			return errors.Wrap(err, "could not marshal the builtin error type")
		}
		code = ErrInternalServerError.StatusCode()
		println("CODE")
		fmt.Printf("%+v\n", j)
		fmt.Printf("%+v\n", err)
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

func Wrap(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := call(w, r, handler); err != nil {
			log.Printf("the rest library encountered an error and was unable to handle it: %s", err)
		}
	}
}
