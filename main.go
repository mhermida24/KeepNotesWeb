package main

import (
	"fmt"
	"log"
	"net/http"

	handlerDB "tpeweb.com/servidor-go/db/handlers"
	sqlc "tpeweb.com/servidor-go/db/sqlc"
	"tpeweb.com/servidor-go/handlers"
)

func main() {
	staticDir := "./static"
	fileServer := http.FileServer(http.Dir(staticDir))
	port := ":8080"
	conn, err := handlerDB.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	queries := sqlc.New(conn)
	userHandler := handlers.NewUserHandler(queries)

	http.Handle("/", fileServer)
	http.HandleFunc("/api/notes", userHandler.NotesHandler)
	http.HandleFunc("/api/notes/", userHandler.NoteHandler)
	http.HandleFunc("/api/folders", userHandler.FoldersHandler)
	http.HandleFunc("/api/folders/", userHandler.FolderHandler)
	http.HandleFunc("/api/users", userHandler.UsersHandler)
	http.HandleFunc("/api/users/", userHandler.SingleUserHandler)
	http.HandleFunc("/api/login", userHandler.LoginHandler)

	fmt.Printf("Servidor ESTÁTICO escuchando en http://localhost%s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
