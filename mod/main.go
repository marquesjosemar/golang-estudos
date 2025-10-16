package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	// <-- Importamos um pacote externo
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá, Mundo com Go 1.25 e Chi!"))
	})

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", r)

}
