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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := Response{
		Success: true,
		Message: msg,
		Data:    data,
	}

	b, err := json.MarshalIndent(res, "", "")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func Created(w http.ResponseWriter, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	res := Response{
		Success: true,
		Message: msg,
		Data:    nil,
	}

	b, _ := json.MarshalIndent(res, "", " ")
	w.Write(b)
}

func BadRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	res := Response{
		Success: false,
		Message: msg,
		Data:    nil,
	}

	b, _ := json.MarshalIndent(res, "", " ")
	w.Write(b)
}

func Unauthorized(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	res := Response{
		Success: false,
		Message: msg,
		Data:    nil,
	}

	b, _ := json.MarshalIndent(res, "", " ")
	w.Write(b)
}

func Forbidden(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	res := Response{
		Success: false,
		Message: msg,
		Data:    nil,
	}

	b, _ := json.MarshalIndent(res, "", " ")
	w.Write(b)
}

func InternalError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	res := Response{
		Success: false,
		Message: msg,
		Data:    nil,
	}

	b, _ := json.MarshalIndent(res, "", " ")
	w.Write(b)
}

func NotFound(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	res := Response{
		Success: false,
		Message: msg,
		Data:    nil,
	}

	b, _ := json.MarshalIndent(res, "", " ")
	w.Write(b)
}
