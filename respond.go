package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func DecodeBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func EncodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func RespondList(w http.ResponseWriter, r *http.Request, status int, data []interface{}) {
	w.Header().Set("Pagination-Count", strconv.Itoa(len(data)))
	Respond(w, r, status, data)
}

func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		EncodeBody(w, r, data)
	}
}

func RespondErr(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
	Respond(w, r, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func RespondHTTPErr(w http.ResponseWriter, r *http.Request, status int) {
	RespondErr(w, r, status, http.StatusText(status))
}
