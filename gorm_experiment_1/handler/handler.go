package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gormex/db"
	"net/http"
	"strconv"
)

func AllNotes(w http.ResponseWriter, r *http.Request) {
	var notes []db.Note

	result := db.DB.Find(&notes)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("CIty", "1")

	err := json.NewEncoder(w).Encode(notes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note db.Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&note)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var note db.Note
	result := db.DB.First(&note, id)
	if result.Error != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result = db.DB.Save(&note)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var note db.Note
	result := db.DB.First(&note, id)
	if result.Error != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	result = db.DB.Delete(&note)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "No note found to delete", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
