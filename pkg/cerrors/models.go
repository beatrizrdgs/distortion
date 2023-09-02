package cerrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CError struct {
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}

func NewCError(message string, statusCode int) CError {
	return CError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func (e CError) Error() string {

	errorStr := ""

	if e.StatusCode != 0 {
		errorStr = fmt.Sprintf("Status Code: %d", e.StatusCode)
	}

	if e.Message != "" {
		errorStr = fmt.Sprintf(" %s", e.Message)
	}

	return errorStr

}

func (e CError) Render(w http.ResponseWriter, r *http.Request) error {

	if e.StatusCode != 0 {
		w.WriteHeader(e.StatusCode)
	}

	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	w.Write(b)

	return nil

}
