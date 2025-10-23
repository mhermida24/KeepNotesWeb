package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

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
func (h *UserHandler) NoteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📌 NoteHandler llamado con método:", r.Method)
	switch r.Method {
	case "GET":
		h.getNoteByID(w, r)
	case "PUT":
		h.updateNote(w, r)
	case "DELETE":
		h.deleteNote(w, r)
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
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdNote)
	if err != nil {
		http.Error(w, "Error al codificar JSON", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) getNoteByID(w http.ResponseWriter, r *http.Request) {
	// Obtener ID desde URL
	idStr := r.URL.Path[len("/notes/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	// Buscar en la base de datos
	note, err := h.queries.GetNote(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, "Error al codificar JSON", http.StatusInternalServerError)
		return
	}

}

func (h *UserHandler) updateNote(w http.ResponseWriter, r *http.Request) {
	// Obtener ID desde URL
	idStr := r.URL.Path[len("/notes/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	// Buscar en la base de datos
	note, err := h.queries.GetNote(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}
	// codigo para updatear nota
	var input struct {
		Title    string  `json:"title"`
		Body     *string `json:"body"`
		FolderID *int32  `json:"folder_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	params := sqlc.UpdateNoteParams{
		ID:    int32(id),
		Title: input.Title,
	}

	if input.Body != nil {
		params.Body = sql.NullString{String: *input.Body, Valid: true}
	} else {
		params.Body = sql.NullString{Valid: false}
	}

	if input.FolderID != nil {
		params.FolderID = sql.NullInt32{Int32: *input.FolderID, Valid: true}
	} else {
		params.FolderID = sql.NullInt32{Valid: false}
	}

	err = h.queries.UpdateNote(r.Context(), params)
	if err != nil {
		http.Error(w, "Error al actualizar la nota", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, "Error al codificar JSON", http.StatusInternalServerError)
		return
	}

}

func (h *UserHandler) deleteNote(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/notes/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
	}
}

func (h *UserHandler) FoldersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📌 NoteHandler llamado con método:", r.Method)
	switch r.Method {
	case "GET":
		h.getFolders(w, r)
	case "PUT":
		h.updateFolder(w, r)
	case "POST":
		h.createFolder(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) FolderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📌 NoteHandler llamado con método:", r.Method)
	switch r.Method {
	case "GET":
		h.getFolderByID(w, r)
	case "PUT":
		h.updateFolder(w, r)
	case "POST":
		h.createFolder(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) getFolders(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	folders, err := h.queries.ListFolders(ctx)
	if err != nil {
		http.Error(w, "Error al listar carpetas: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}

func (h *UserHandler) createFolder(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var folder struct {
		Name           string  `json:"name"`
		Description    *string `json:"description"`
		ParentFolderID *int32  `json:"parent_folder_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&folder); err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if folder.Name == "" {
		http.Error(w, "El nombre es obligatorio", http.StatusBadRequest)
		return
	}

	params := sqlc.CreateFolderParams{
		Name: folder.Name,
	}

	if folder.Description != nil {
		params.Description = sql.NullString{String: *folder.Description, Valid: true}
	} else {
		params.Description = sql.NullString{Valid: false}
	}

	if folder.ParentFolderID != nil {
		params.ParentFolderID = sql.NullInt32{Int32: *folder.ParentFolderID, Valid: true}
	} else {
		params.ParentFolderID = sql.NullInt32{Valid: false}
	}

	createdFolder, err := h.queries.CreateFolder(ctx, params)
	if err != nil {
		http.Error(w, "Error al crear carpeta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdFolder)
	if err != nil {
		http.Error(w, "Error al codificar JSON", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) getFolderByID(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) updateFolder(w http.ResponseWriter, r *http.Request) {

}
