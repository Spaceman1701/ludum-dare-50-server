package web

import (
	"fmt"
	"net/http"

	"github.com/Spaceman1701/ludum-dare-50-server/persistence"
)

type handlerFunc = func(w http.ResponseWriter, req *http.Request)

func AddEntry(db persistence.DB) handlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "add entry")
		w.WriteHeader(100)
	}
}

func GetEntrySummary(db persistence.DB) handlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "get entries")
		w.WriteHeader(100)
	}
}

func RunServer(addr string, db persistence.DB) error {
	http.HandleFunc("/add_entry", AddEntry(db))
	http.HandleFunc("/get_entries", GetEntrySummary(db))
	return http.ListenAndServe(addr, nil)
}
