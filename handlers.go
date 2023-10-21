package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bicycle-Guild/kaktys/models"
	"github.com/Bicycle-Guild/kaktys/repository"
	"github.com/gorilla/mux"
)

func create_subject_http(w http.ResponseWriter, r *http.Request) {
	var new_subject models.CreateSubject

	json.NewDecoder(r.Body).Decode(&new_subject)

	json.NewEncoder(w).Encode(repository.Create_subject(new_subject))
}

func get_subjects_http(w http.ResponseWriter, r *http.Request) {
	subjects := repository.Get_subjects()

	if subjects == nil {
		subjects = []models.Subject{}
	}

	json.NewEncoder(w).Encode(subjects)
}

func get_subject_http(w http.ResponseWriter, r *http.Request) {
	subject, err := repository.Get_Subject(mux.Vars(r)["subjectID"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Not Found"})
		return
	}

	json.NewEncoder(w).Encode(subject)
}
