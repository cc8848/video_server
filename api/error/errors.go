package error

import "net/http"

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type Response struct {
	HttpSC int
	Error  Err
}

var (
	RequestBodyParseFailedError = Response{
		HttpSC: http.StatusBadRequest,
		Error: Err{
			Error:     "Request body is not correct.",
			ErrorCode: "001",
		},
	}
	NotAuthUserError = Response{
		HttpSC: http.StatusUnauthorized,
		Error: Err{
			Error:     "User authentication failed.",
			ErrorCode: "002",
		},
	}
)
