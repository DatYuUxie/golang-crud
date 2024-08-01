package dao

import (
	"CRUDProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateSubjectHandler(w http.ResponseWriter, r *http.Request) {
	var subject models.Subject

	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.CreateSubject(&subject); err != nil {
		http.Error(w, "Error creating subject: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subject)
}

func GetSubjectByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format: "+err.Error(), http.StatusBadRequest)
		return
	}

	subject, err := models.GetSubjectByID(uint(id))
	if err != nil {
		http.Error(w, "Subject not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subject)
}

func GetSubjectsHandler(w http.ResponseWriter, r *http.Request) {
	subjects, err := models.GetAllSubjects()
	if err != nil {
		http.Error(w, "Error retrieving subjects: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subjects)
}

func UpdateSubjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format: "+err.Error(), http.StatusBadRequest)
		return
	}

	var subject models.Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.UpdateSubjectByID(uint(id), &subject); err != nil {
		http.Error(w, "Error updating subject: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteSubjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.DeleteSubject(uint(id)); err != nil {
		http.Error(w, "Error deleting subject: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
