package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func json_middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		f(w, r)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/subjects", json_middleware(get_subjects_http)).Methods("GET")
	r.HandleFunc("/api/subjects/{subjectID}", json_middleware(get_subject_http)).Methods("GET")
	r.HandleFunc("/api/subjects", json_middleware(create_subject_http)).Methods("POST")

	http.ListenAndServe(":8080", r)
}
