package response

import "net/http"

type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type Err struct {
	Error     string `json:"response"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

var (
	RequestBodyParseFailedError = ErrorResponse{
		HttpSC: http.StatusBadRequest,
		Error: Err{
			Error:     "Request body is not correct.",
			ErrorCode: "001",
		},
	}
	NoAuthUserError = ErrorResponse{
		HttpSC: http.StatusUnauthorized,
		Error: Err{
			Error:     "User authentication failed.",
			ErrorCode: "002",
		},
	}
	DBError = ErrorResponse{
		HttpSC: http.StatusInternalServerError,
		Error: Err{
			Error:     "DB ops  failed",
			ErrorCode: "003",
		},
	}
	InternalFaults = ErrorResponse{
		HttpSC: http.StatusInternalServerError,
		Error: Err{
			Error:     "Internal Service error",
			ErrorCode: "004",
		},
	}
)
