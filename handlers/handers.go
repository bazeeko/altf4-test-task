package handlers

import (
	"app/binance"
	"app/ws"
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	book, err := binance.GetBook()
	if err != nil {
		log.Fatalln("Couldn't get books", err)
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln("Coudn't parse files", err)
	}

	t.Execute(w, book)
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	websocket, err := ws.Upgrade(w, r)
	if err != nil {
		log.Fatalln("Couldn't connect websocket", err)
	}

	go ws.Writer(websocket)
}
