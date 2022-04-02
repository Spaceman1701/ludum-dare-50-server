package web

import (
	"fmt"
	"net/http"
)

type handlerFunc = func(w http.ResponseWriter, req *http.Request)

func AddEntry() handlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "add entry")
		w.WriteHeader(100)
	}
}

func GetEntrySummary() handlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "get entries")
		w.WriteHeader(100)
	}
}

func RunServer(addr string) error {
	http.HandleFunc("/add_entry", AddEntry())
	http.HandleFunc("/get_entries", GetEntrySummary())
	return http.ListenAndServe(addr, nil)
}
