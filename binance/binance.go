package binance

import (
	"encoding/json"
	"net/http"
)

type BID []string

type ASK []string

type Book struct {
	LastUpdateId uint64 `json:"lastUpdateId"`
	BIDs         []BID  `json:"bids"`
	ASKs         []ASK  `json:"asks"`
}

func getJSON(url string, book *Book) error {
	r, err := (&http.Client{}).Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(book)
}

func (book *Book) getFirst15() {
	book.BIDs = book.BIDs[:15]
	book.ASKs = book.ASKs[:15]
}

func GetBook() (*Book, error) {
	book := new(Book)

	if err := getJSON("https://api1.binance.com/api/v3/depth?symbol=BTCUSDT&limit=20", book); err != nil {
		return book, err
	}

	book.getFirst15()

	return book, nil
}
