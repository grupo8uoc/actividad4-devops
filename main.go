package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time" // IMPORTANTE: Nuevo import necesario
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<html>
			<head><title>Actividad 4</title></head>
			<body style="text-align:center; font-family: sans-serif;">
				<h1>Sistema de Despliegue Seguro</h1>
				<p>Implementación por el Grupo - VERSION SEGURA</p>
				<img src="/static/logo.png" alt="Logo Grupo" width="300">
			</body>
		</html>`
		fmt.Fprint(w, html)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// CONFIGURACIÓN SEGURA DEL SERVIDOR (Corrige el fallo G114 de Gosec)
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      nil,              // Usa el DefaultServeMux
		ReadTimeout:  10 * time.Second, // Timeout para leer petición
		WriteTimeout: 10 * time.Second, // Timeout para escribir respuesta
		IdleTimeout:  15 * time.Second, // Timeout conexión inactiva
	}

	log.Printf("Iniciando servidor SEGURO en puerto %s...", port)
	// Usamos server.ListenAndServe en vez de http.ListenAndServe
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
