# rest [![GoDoc](https://godoc.org/github.com/boreq/rest?status.svg)](https://godoc.org/github.com/boreq/rest)

Package rest implements a framework for creating a JSON REST API.

It automatically serializes the responses as JSON and provides the user with
the ability to implement a clearer flow of control within the handler thanks
to encapsulating all calls to the response writer within a single return
value from the handler.

    func handler(r *http.Request) rest.RestResponse {
        value, err := someFunction()
        if err != nil {
            return rest.InternalServerError
        }
        return rest.NewResponse(value)
    }