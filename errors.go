package rest

import "net/http"

type errorBody struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// Error represents an error response from the API. Returning it from your
// handler will produce a message serialized in the following way:
//
//  {
//  	"statusCode": 123,
//  	"message": "Provided message."
//  }
//
// Error encapsulates the response which means that all methods available on
// that type can be used freely on an error.
type Error struct {
	Response
}

// NewError creates a new error with the specified status code and message.
func NewError(statusCode int, message string) Error {
	body := errorBody{
		StatusCode: statusCode,
		Message:    message,
	}
	return Error{
		Response: NewResponse(body).WithStatusCode(statusCode),
	}
}

// WithMessage returns a new error with the changed message. This method makes
// it easy to use the provided builtin error types with non-standard messages.
func (e Error) WithMessage(message string) Error {
	r2 := e.clone()
	r2.body = errorBody{
		StatusCode: r2.StatusCode(),
		Message:    message,
	}
	return Error{r2}
}

// Predefined 4xx client errors.
var (
	ErrBadRequest                   = NewError(http.StatusBadRequest, "Bad request.")
	ErrUnauthorized                 = NewError(http.StatusUnauthorized, "Unauthorized.")
	ErrPaymentRequired              = NewError(http.StatusPaymentRequired, "Payment required.")
	ErrForbidden                    = NewError(http.StatusForbidden, "Forbidden.")
	ErrNotFound                     = NewError(http.StatusForbidden, "Not found.")
	ErrMethodNotAllowed             = NewError(http.StatusMethodNotAllowed, "Method not allowed.")
	ErrNotAcceptable                = NewError(http.StatusNotAcceptable, "Not acceptable.")
	ErrProxyAuthRequired            = NewError(http.StatusProxyAuthRequired, "Proxy auth required.")
	ErrRequestTimeout               = NewError(http.StatusRequestTimeout, "Request timeout.")
	ErrConflict                     = NewError(http.StatusConflict, "Conflict.")
	ErrGone                         = NewError(http.StatusGone, "Gone.")
	ErrLengthRequired               = NewError(http.StatusLengthRequired, "Length required.")
	ErrPreconditionFailed           = NewError(http.StatusPreconditionFailed, "Precondition failed.")
	ErrRequestEntityTooLarge        = NewError(http.StatusRequestEntityTooLarge, "Request entity too large.")
	ErrRequestURITooLong            = NewError(http.StatusRequestURITooLong, "Request URI too long.")
	ErrUnsupportedMediaType         = NewError(http.StatusUnsupportedMediaType, "Unsupported media type.")
	ErrRequestedRangeNotSatisfiable = NewError(http.StatusRequestedRangeNotSatisfiable, "Requested range not satisfiable.")
	ErrExpectationFailed            = NewError(http.StatusExpectationFailed, "Expectation failed.")
	ErrTeapot                       = NewError(http.StatusTeapot, "I'm a teapot.")
	ErrMisdirectedRequest           = NewError(http.StatusMisdirectedRequest, "Misredirected request.")
	ErrUnprocessableEntity          = NewError(http.StatusUnprocessableEntity, "Unprocessable entity.")
	ErrLocked                       = NewError(http.StatusLocked, "Locked.")
	ErrFailedDependency             = NewError(http.StatusFailedDependency, "Failed dependency.")
	ErrTooEarly                     = NewError(http.StatusTooEarly, "Too early.")
	ErrUpgradeRequired              = NewError(http.StatusUpgradeRequired, "Upgrade required.")
	ErrPreconditionRequired         = NewError(http.StatusPreconditionRequired, "Precondition required.")
	ErrTooManyRequests              = NewError(http.StatusTooManyRequests, "Too many requests.")
	ErrRequestHeaderFieldsTooLarge  = NewError(http.StatusRequestHeaderFieldsTooLarge, "Request header fields too large.")
	ErrUnavailableForLegalReasons   = NewError(http.StatusUnavailableForLegalReasons, "Unavailable for legal reasons.")
)

// Predefined 5xx server errors.
var (
	ErrInternalServerError           = NewError(http.StatusInternalServerError, "Internal server error.")
	ErrNotImplemented                = NewError(http.StatusNotImplemented, "Not implemented.")
	ErrBadGateway                    = NewError(http.StatusBadGateway, "Bad gateway.")
	ErrServiceUnavailable            = NewError(http.StatusServiceUnavailable, "Service unavailable.")
	ErrGatewayTimeout                = NewError(http.StatusGatewayTimeout, "Gateway timeout.")
	ErrHTTPVersionNotSupported       = NewError(http.StatusHTTPVersionNotSupported, "HTTP version not supported.")
	ErrVariantAlsoNegotiates         = NewError(http.StatusVariantAlsoNegotiates, "Variant also negotiates.")
	ErrInsufficientStorage           = NewError(http.StatusInsufficientStorage, "Insufficient storage.")
	ErrLoopDetected                  = NewError(http.StatusLoopDetected, "Loop detected.")
	ErrNotExtended                   = NewError(http.StatusNotExtended, "Not extended.")
	ErrNetworkAuthenticationRequired = NewError(http.StatusNetworkAuthenticationRequired, "Network authentication required.")
)
