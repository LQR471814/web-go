package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("../../public")))
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}
