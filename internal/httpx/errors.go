package httpx

import (
	"encoding/json"
	"net/http"
)

type Code struct {
	value string
}

// Predefined error codes.
var (
	CodeInvalidID = Code{
		value: "invalid_id",
	}

	CodeInternalError = Code{
		value: "internal_error",
	}

	CodeNotFound = Code{
		value: "not_found",
	}

	CodeMalformedJSON = Code{
		value: "malformed_json",
	}

	CodeValidationFailed = Code{
		value: "validation_failed",
	}

	CodeUnauthenticated = Code{
		value: "unauthenticated",
	}

	CodeForbidden = Code{
		value: "forbidden",
	}

	CodeConflict = Code{
		value: "conflict",
	}

	CodeRateLimited = Code{
		value: "rate_limited",
	}
)

// MarshalJSON makes Code serialize as a JSON string.
//
// Example:
//
//	CodeInternalError -> "internal_error"
func (c Code) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

type errorEnvelope struct {
	Error errorPayload `json:"error"`
}

type errorPayload struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

// Error writes a JSON error response.
func Error(w http.ResponseWriter, status int, message string, code Code) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(errorEnvelope{
		Error: errorPayload{
			Code:    code,
			Message: message,
		},
	})
}

// package httpx

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type Code string

// const (
// 	CodeInvalidID        Code = "invalid_id"
// 	CodeInternalError    Code = "internal_error"
// 	CodeNotFound         Code = "not_found"
// 	CodeMalFormedJSON    Code = "malformed_json"
// 	CodeValidationFailed Code = "validation_failed"
// 	CodeUnauthenticated  Code = "unauthenticated"
// 	CodeForbidden        Code = "forbidden"
// 	CodeConflict         Code = "conflict"
// 	CodeRateLimited      Code = "rate_limited"
// )

// type errorEnvelope struct {
// 	Error errorPayload `json:"error"`
// }

// type errorPayload struct {
// 	Code    Code   `json:"code"`
// 	Message string `json:"message"`
// }

// func Error(w http.ResponseWriter, status int, message string, code Code) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)

// 	_ = json.NewEncoder(w).Encode(errorEnvelope{Error: errorPayload{
// 		Code:    code,
// 		Message: message,
// 	}})
// }
