package main

import (
	"app/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/ws", handlers.WsHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
