package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type AcceptedResponse struct{}

type CreatedResponse struct{}

type OkResponse struct {
	StatusText string `json:"status,omitempty"`
}

var env = os.Getenv("ENV")

type IDResponse struct {
	ID string `json:"id"`
}

type ErrResponse struct {
	StatusText string `json:"status"`
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// render renders a single payload and respond to the client request.
func Render[T any](w http.ResponseWriter, r *http.Request, status int, input T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(input); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func ErrInvalidRequest(err error, w http.ResponseWriter, r *http.Request) {
	errMsg := "error"
	if env == "dev" {
		errMsg = err.Error()
	}

	Render(w, r, http.StatusBadRequest, ErrResponse{
		StatusText: "Invalid request",
		ErrorText:  errMsg,
	})
}

func ErrNotFound(w http.ResponseWriter, r *http.Request) {
	Render(w, r, http.StatusNotFound, ErrResponse{
		StatusText: "Not found",
	})
}

func ErrInternalServer(err error, w http.ResponseWriter, r *http.Request) {
	errMsg := "error"
	if env == "dev" {
		errMsg = err.Error()
	}

	Render(w, r, http.StatusInternalServerError, ErrResponse{
		StatusText: "Internal server error",
		ErrorText:  errMsg,
	})
}

func ErrUnauthorizedRequest(err error, w http.ResponseWriter, r *http.Request) {
	errMsg := "error"
	if env == "dev" {
		errMsg = err.Error()
	}

	Render(w, r, http.StatusForbidden, ErrResponse{
		StatusText: "Unauthorized",
		ErrorText:  errMsg,
	})
}

func Accepted(w http.ResponseWriter, r *http.Request) {
	Render(w, r, http.StatusAccepted, AcceptedResponse{})
}

func Ok(w http.ResponseWriter, r *http.Request) {
	Render(w, r, http.StatusOK, OkResponse{
		StatusText: "ok",
	})
}
