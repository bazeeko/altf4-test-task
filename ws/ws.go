package ws

import (
	"app/binance"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return upgrader.Upgrade(w, r, nil)
}

func Writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(1 * time.Second)

		for t := range ticker.C {
			fmt.Printf("Upgrading: %v\n", t)

			book, err := binance.GetBook()
			if err != nil {
				log.Println(err)
			}

			jsonString, err := json.Marshal(book)
			if err != nil {
				log.Println(err)
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				log.Println(err)
			}
		}
	}
}
