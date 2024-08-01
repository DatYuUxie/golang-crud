package dao

import (
	"CRUDProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateClassHandler(w http.ResponseWriter, r *http.Request) {
	var class models.Class

	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.CreateClass(&class); err != nil {
		http.Error(w, "Error creating class: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

func GetClassByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format: "+err.Error(), http.StatusBadRequest)
		return
	}

	class, err := models.GetClassByID(uint(id))
	if err != nil {
		http.Error(w, "Class not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}

func GetClassesHandler(w http.ResponseWriter, r *http.Request) {
	classes, err := models.GetAllClasses()
	if err != nil {
		http.Error(w, "Error retrieving classes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

func UpdateClassHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format: "+err.Error(), http.StatusBadRequest)
		return
	}

	var class models.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.UpdateClass(uint(id), &class); err != nil {
		http.Error(w, "Error updating class: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteClassHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.DeleteClass(uint(id)); err != nil {
		http.Error(w, "Error deleting class: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func RegisterClassHandler(w http.ResponseWriter, r *http.Request) {
	classIDStr := r.URL.Query().Get("class_id")
	studentIDStr := r.URL.Query().Get("student_id")

	classID, err := strconv.ParseUint(classIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Class ID format", http.StatusBadRequest)
		return
	}

	studentID, err := strconv.ParseUint(studentIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Student ID format", http.StatusBadRequest)
		return
	}

	if err := models.RegisterClass(uint(classID), uint(studentID)); err != nil {
		http.Error(w, "Error registering student to class: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
