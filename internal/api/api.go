package api

import (
    "encoding/json"
    "net/http"
)

// Payload represents the data and metadata of the API response.
type Payload struct {
    Meta json.RawMessage `json:"meta,omitempty"`
    Data json.RawMessage `json:"data,omitempty"`
}

// Response represents the structure of the API response.
type Response struct {
    Payload   Payload  `json:"payload"`
    Messages  []string `json:"messages"`
    ErrorCode int      `json:"errorCode"`
}

// ErrorResponse represents the structure of an error response.
type ErrorResponse struct {
    Error     string `json:"error"`
    ErrorCode int    `json:"errorCode"`
}

// NewResponse creates a new API response with the given data and messages.
func NewResponse(data interface{}, messages []string) (*Response, error) {
    dataBytes, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    return &Response{
        Payload: Payload{
            Data: dataBytes,
        },
        Messages:  messages,
        ErrorCode: 0,
    }, nil
}

// NewErrorResponse creates a new error response with the given error message and code.
func NewErrorResponse(err error, errorCode int) *ErrorResponse {
    return &ErrorResponse{
        Error:     err.Error(),
        ErrorCode: errorCode,
    }
}

// WriteResponse writes the API response to the HTTP response writer.
func WriteResponse(w http.ResponseWriter, data interface{}, statusCode int) {
    response, err := NewResponse(data, nil)
    if err != nil {
        WriteErrorResponse(w, err, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(response)
}

// WriteErrorResponse writes the error response to the HTTP response writer.
func WriteErrorResponse(w http.ResponseWriter, err error, errorCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(NewErrorResponse(err, errorCode))
}
