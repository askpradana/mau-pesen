package common

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(w http.ResponseWriter, msg string, data interface{}) {
	json.NewEncoder(w).Encode(Response{true, msg, data})
}

func Created(w http.ResponseWriter, msg string, data interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{true, msg, data})
}

func BadRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Response{false, msg, nil})
}

func Unauthorized(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(Response{false, msg, nil})
}

func Forbidden(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(Response{false, msg, nil})
}

func InternalError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Response{false, msg, nil})
}

func NotFound(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{true, msg, nil})
}
