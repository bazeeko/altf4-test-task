package main

import (
	"app/handlers"
	"log"
	"net/http"
)

func main() {
	files := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/ws", handlers.WsHandler)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
