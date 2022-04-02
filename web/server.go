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
		fmt.Fprintf(w, "%v", history)
		db.WriteEntry(history)
	}
}

func GetEntrySummary(db persistence.DB) handlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "get entries\n")

		history, _ := db.ReadAllEntries()
		for i, h := range history {
			fmt.Fprintf(w, "entry %v: %v\n", i, h)
		}
	}
}

func RunServer(addr string, db persistence.DB) error {
	http.HandleFunc("/add_history", AddEntry(db))
	http.HandleFunc("/get_game_state", GetEntrySummary(db))
	return http.ListenAndServe(addr, nil)
}
