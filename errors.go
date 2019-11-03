package rest

type errorBody struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type Error struct {
	Response
}

func NewError(statusCode int, message string) Error {
	body := errorBody{
		StatusCode: statusCode,
		Message:    message,
	}
	return Error{
		Response: NewResponse(body).WithStatusCode(statusCode),
	}
}

func (e Error) WithMessage(message string) Error {
	r2 := e.clone()
	r2.body = errorBody{
		StatusCode: r2.StatusCode(),
		Message:    message,
	}
	return Error{r2}
}

var ErrBadRequest = NewError(400, "Bad request.")
var ErrUnauthorized = NewError(401, "Unauthorized.")
var ErrPaymentRequired = NewError(402, "Payment required.")
var ErrForbidden = NewError(403, "Forbidden.")
var ErrNotFound = NewError(404, "Not found.")
var ErrMethodNotAllowed = NewError(405, "Method not allowed.")
var ErrNotAcceptable = NewError(406, "Not acceptable.")
var ErrProxyAuthRequired = NewError(407, "Proxy auth required.")
var ErrRequestTimeout = NewError(408, "Request timeout.")
var ErrConflict = NewError(409, "Conflict.")
var ErrGone = NewError(410, "Gone.")
var ErrLengthRequired = NewError(411, "Length required.")
var ErrPreconditionFailed = NewError(412, "Precondition failed.")
var ErrRequestEntityTooLarge = NewError(413, "Request entity too large.")
var ErrRequestURITooLong = NewError(414, "Request URI too long.")
var ErrUnsupportedMediaType = NewError(415, "Unsupported media type.")
var ErrRequestedRangeNotSatisfiable = NewError(416, "Requested range not satisfiable.")
var ErrExpectationFailed = NewError(417, "Expectation failed.")
var ErrTeapot = NewError(418, "I'm a teapot.")
var ErrMisdirectedRequest = NewError(421, "Misredirected request.")
var ErrUnprocessableEntity = NewError(422, "Unprocessable entity.")
var ErrLocked = NewError(423, "Locked.")
var ErrFailedDependency = NewError(424, "Failed dependency.")
var ErrTooEarly = NewError(425, "Too early.")
var ErrUpgradeRequired = NewError(426, "Upgrade required.")
var ErrPreconditionRequired = NewError(428, "Precondition required.")
var ErrTooManyRequests = NewError(429, "Too many requests.")
var ErrRequestHeaderFieldsTooLarge = NewError(431, "Request header fields too large.")
var ErrUnavailableForLegalReasons = NewError(451, "Unavailable for legal reasons.")

var ErrInternalServerError = NewError(500, "Internal server error.")
var ErrNotImplemented = NewError(501, "Not implemented.")
var ErrBadGateway = NewError(502, "Bad gateway.")
var ErrServiceUnavailable = NewError(503, "Service unavailable.")
var ErrGatewayTimeout = NewError(504, "Gateway timeout.")
var ErrHTTPVersionNotSupported = NewError(505, "HTTP version not supported.")
var ErrVariantAlsoNegotiates = NewError(506, "Variant also negotiates.")
var ErrInsufficientStorage = NewError(507, "Insufficient storage.")
var ErrLoopDetected = NewError(508, "Loop detected.")
var ErrNotExtended = NewError(510, "Not extended.")
var ErrNetworkAuthenticationRequired = NewError(511, "Network authentication required.")
