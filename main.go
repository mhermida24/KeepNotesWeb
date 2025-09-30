package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	sqlc "tpeweb.com/servidor-go/db/sqlc"
)

func main() {
	staticDir := "./static"
	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fileServer)
	port := ":8080"
	fmt.Printf("Servidor ESTÁTICO escuchando en http://localhost%s\n", port)
	fmt.Printf("Sirviendo archivos desde: %s\n", staticDir)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
