package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Spaceman1701/ludum-dare-50-server/model"
	"github.com/Spaceman1701/ludum-dare-50-server/persistence"
)

type handlerFunc = func(w http.ResponseWriter, req *http.Request)

func AddEntry(db persistence.DB) handlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var history model.GameHistory
		err := json.NewDecoder(req.Body).Decode(&history)
		if err != nil {
			fmt.Fprintf(w, "bad\n")
			return
		}
		fmt.Fprintf(w, "add entry")
	}
}

func GetEntrySummary(db persistence.DB) handlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "get entries")
	}
}

func RunServer(addr string, db persistence.DB) error {
	http.HandleFunc("/add_entry", AddEntry(db))
	http.HandleFunc("/get_entries", GetEntrySummary(db))
	return http.ListenAndServe(addr, nil)
}
