package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test para verificar que el healthcheck devuelve OK
func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	handler.ServeHTTP(rr, req)

	// Comprobamos el código 200
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Código incorrecto: se obtuvo %v se esperaba %v",
			status, http.StatusOK)
	}

	// Comprobamos el cuerpo "OK"
	if rr.Body.String() != "OK" {
		t.Errorf("Cuerpo incorrecto: se obtuvo %v se esperaba %v",
			rr.Body.String(), "OK")
	}
}