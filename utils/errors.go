package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/whatcrm/go-wildberries/models"
)

type APIError struct {
	StatusCode int

	Title      string  `json:"title"`
	Detail     string  `json:"detail"`
	Code       string  `json:"code"`
	RequestID  string  `json:"requestId"`
	Origin     string  `json:"origin"`
	Status     float64 `json:"status"`
	StatusText string  `json:"statusText"`
	Timestamp  string  `json:"timestamp"`

	RawBody    []byte
	RetryAfter string
}

func (e *APIError) Error() string {
	if e.Detail != "" {
		return e.Detail
	}
	if e.Title != "" {
		return e.Title
	}
	return "WILDBERRIES API ERROR"
}

func ParseAPIError(resp *http.Response) error {
	raw, _ := io.ReadAll(resp.Body)

	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		RawBody:    raw,
		RetryAfter: resp.Header.Get("Retry-After"),
	}

	_ = json.Unmarshal(raw, apiErr)

	// Some methods return another error schema without code/statusText/timestamp.
	if apiErr.Title == "" && apiErr.Detail == "" {
		var errBody models.ErrorResponse
		if err := json.Unmarshal(raw, &errBody); err == nil {
			apiErr.Title = errBody.Title
			apiErr.Detail = errBody.Detail
			apiErr.RequestID = errBody.RequestID
			apiErr.Origin = errBody.Origin
			apiErr.Status = errBody.Status
		}
	}

	if apiErr.Detail == "" {
		apiErr.Detail = strings.TrimSpace(string(raw))
	}

	return apiErr
}
