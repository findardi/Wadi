package response

import (
	"encoding/json"
	"net/http"
)

type Envelope struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func writeJson(w http.ResponseWriter, status int, env Envelope) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(env)
}

func Success(w http.ResponseWriter, status int, message string, data any) {
	writeJson(w, status, Envelope{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SuccessWithMeta(w http.ResponseWriter, status int, messgae string, data any, meta Meta) {
	writeJson(w, status, Envelope{
		Success: true,
		Message: messgae,
		Data:    data,
		Meta:    &meta,
	})
}

func Error(w http.ResponseWriter, status int, message string, errs any) {
	writeJson(w, status, Envelope{
		Success: false,
		Message: message,
		Errors:  errs,
	})
}
