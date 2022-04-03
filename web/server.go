package web

import (
	"net/http"

	"github.com/Spaceman1701/ludum-dare-50-server/persistence"
)

type handlerFunc = func(w http.ResponseWriter, req *http.Request)

func AddEntry(db *persistence.DB) handlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func GetEntrySummary(db *persistence.DB) handlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func RunServer(addr string, db *persistence.DB) error {
	http.HandleFunc("/add_history", AddEntry(db))
	http.HandleFunc("/get_game_state", GetEntrySummary(db))
	return http.ListenAndServe(addr, nil)
}
