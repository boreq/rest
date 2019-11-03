package rest_test

import (
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boreq/rest"
	"github.com/stretchr/testify/require"
)

func TestResponse(t *testing.T) {
	handler := func(r *http.Request) rest.RestResponse {
		responseBody := someBody()
		return rest.NewResponse(responseBody)
	}
	res, body := execute(t, handler)
	require.Equal(t, `{"field":"value"}`, string(body))
	require.Equal(t, 200, res.StatusCode)
	require.Equal(t, "application/json", res.Header.Get("Content-Type"))
}

func TestResponseCustomContentType(t *testing.T) {
	handler := func(r *http.Request) rest.RestResponse {
		responseBody := someBody()
		return rest.NewResponse(responseBody).WithHeader("Content-Type", "application/xml")
	}
	res, body := execute(t, handler)
	require.Equal(t, `{"field":"value"}`, string(body))
	require.Equal(t, 200, res.StatusCode)
	require.Equal(t, "application/xml", res.Header.Get("Content-Type"))
	require.Equal(t, 1, len(res.Header["Content-Type"]))
}

func TestResponseWithCode(t *testing.T) {
	handler := func(r *http.Request) rest.RestResponse {
		responseBody := someBody()
		return rest.NewResponse(responseBody).WithStatusCode(202)
	}
	res, body := execute(t, handler)
	require.Equal(t, `{"field":"value"}`, string(body))
	require.Equal(t, 202, res.StatusCode)
	require.Equal(t, "application/json", res.Header.Get("Content-Type"))
}

func TestResponseWithHeader(t *testing.T) {
	handler := func(r *http.Request) rest.RestResponse {
		responseBody := someBody()
		return rest.NewResponse(responseBody).WithHeader("Header", "value")
	}
	res, body := execute(t, handler)
	require.Equal(t, `{"field":"value"}`, string(body))
	require.Equal(t, 200, res.StatusCode)
	require.Equal(t, "application/json", res.Header.Get("Content-Type"))
	require.Equal(t, "value", res.Header.Get("Header"))
}

func TestResponseImpossilbeToMarshal(t *testing.T) {
	handler := func(r *http.Request) rest.RestResponse {
		return rest.NewResponse(math.Inf(1))
	}
	res, body := execute(t, handler)
	require.Equal(t, `{"statusCode":500,"message":"Internal server error."}`, string(body))
	require.Equal(t, 500, res.StatusCode)
	require.Equal(t, "application/json", res.Header.Get("Content-Type"))
}

func TestError(t *testing.T) {
	handler := func(r *http.Request) rest.RestResponse {
		return rest.ErrNotAcceptable
	}
	res, body := execute(t, handler)
	require.Equal(t, `{"statusCode":406,"message":"Not acceptable."}`, string(body))
	require.Equal(t, 406, res.StatusCode)
	require.Equal(t, "application/json", res.Header.Get("Content-Type"))
}

func TestErrorWithMessage(t *testing.T) {
	handler := func(r *http.Request) rest.RestResponse {
		return rest.ErrNotAcceptable.WithMessage("Custom message.")
	}
	res, body := execute(t, handler)
	require.Equal(t, `{"statusCode":406,"message":"Custom message."}`, string(body))
	require.Equal(t, 406, res.StatusCode)
	require.Equal(t, "application/json", res.Header.Get("Content-Type"))
}

func execute(t *testing.T, handler rest.HandlerFunc) (*http.Response, []byte) {
	server := httptest.NewServer(rest.Wrap(handler))
	defer server.Close()

	res, err := http.Get(server.URL)
	require.NoError(t, err)

	body, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)

	return res, body
}

func someBody() interface{} {
	return struct {
		Field string `json:"field"`
	}{
		Field: "value",
	}
}

func ExampleWrap() {
	handler := func(r *http.Request) rest.RestResponse {
		return rest.NewResponse("response").WithHeader("X-Clacks-Overhead", "GNU Terry Pratchett")
	}

	server := httptest.NewServer(rest.Wrap(handler))
	defer server.Close()
}

func ExampleError_WithMessage() {
	handler := func(r *http.Request) rest.RestResponse {
		return rest.ErrInternalServerError.WithMessage("Custom error message.")
	}

	server := httptest.NewServer(rest.Wrap(handler))
	defer server.Close()
}
