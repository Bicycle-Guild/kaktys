package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Test")
	})

	http.ListenAndServe(":8080", r)
}
