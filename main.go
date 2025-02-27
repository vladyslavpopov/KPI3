package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)
	log.Println("Server started on :8795")
	if err := http.ListenAndServe(":8795", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
