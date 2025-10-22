package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	sqlc "tpeweb.com/servidor-go/db/sqlc"
)

type UserHandler struct {
	queries *sqlc.Queries
}

func NewUserHandler(q *sqlc.Queries) *UserHandler {
	return &UserHandler{queries: q}
}

func (h *UserHandler) NotesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📌 NotesHandler llamado con método:", r.Method)
	switch r.Method {
	case "GET":
		h.getNotes(w, r)
	case "POST":
		h.createNote(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) getNotes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	notes, err := h.queries.ListNotes(ctx)
	if err != nil {
		http.Error(w, "Error al listar notas: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func (h *UserHandler) createNote(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	var note struct {
		Title    string  `json:"title"`
		Body     *string `json:"body"`      // puntero para distinguir NULL
		FolderID *int32  `json:"folder_id"` // puntero para distinguir NULL
	}

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if note.Title == "" {
		http.Error(w, "El título es obligatorio", http.StatusBadRequest)
		return
	}

	// Preparar parámetros para sqlc
	params := sqlc.CreateNoteParams{
		Title: note.Title,
	}

	if note.Body != nil {
		params.Body = sql.NullString{String: *note.Body, Valid: true}
	} else {
		params.Body = sql.NullString{Valid: false}
	}

	if note.FolderID != nil {
		params.FolderID = sql.NullInt32{Int32: *note.FolderID, Valid: true}
	} else {
		params.FolderID = sql.NullInt32{Valid: false}
	}

	// Crear la nota en la base de datos
	createdNote, err := h.queries.CreateNote(ctx, params)
	if err != nil {
		http.Error(w, "Error al crear nota: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdNote)
}
