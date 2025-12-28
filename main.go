package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Definimos el puerto. Usamos variable de entorno PORT o 8080 por defecto.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Servidor de archivos estáticos (busca carpeta 'static')
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<html>
			<head><title>Actividad 4</title></head>
			<body style="text-align:center; font-family: sans-serif;">
				<h1>Sistema de Despliegue Seguro</h1>
				<p>Implementación por el Grupo</p>
				<img src="/static/logo.png" alt="Logo Grupo" width="300">
			</body>
		</html>`
		fmt.Fprint(w, html)
	})

	// Health check (vital para sistemas de despliegue modernos)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("Iniciando servidor en puerto %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}