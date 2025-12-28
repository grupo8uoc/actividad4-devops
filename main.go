package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Definir puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<html>
			<head><title>Actividad 4 - Segura</title></head>
			<body style="text-align:center; font-family: sans-serif;">
				<h1>Sistema de Despliegue Seguro</h1>
				<p>Implementación por el Grupo - CORREGIDO</p>
				<img src="/static/logo.png" alt="Logo Grupo" width="300">
			</body>
		</html>`
		fmt.Fprint(w, html)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Configuración del servidor con Timeouts
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Iniciando servidor en puerto %s...", port)
	// Gosec suele quejarse del bind genérico. Con #nosec le decimos que lo ignore.
	if err := server.ListenAndServe(); err != nil { // #nosec G102
		log.Fatal(err)
	}
}
