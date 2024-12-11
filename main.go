package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estructura para manejar el webhook
type WebhookPayload struct {
	Message string `json:"message"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload WebhookPayload

	// Decodificar el JSON recibido
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Mostrar el mensaje recibido
	fmt.Fprintf(w, "Webhook received: %s", payload.Message)
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	// Iniciar servidor
	log.Fatal(http.ListenAndServe(":8080", nil))
}