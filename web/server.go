package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Spaceman1701/ludum-dare-50-server/model"
	"github.com/Spaceman1701/ludum-dare-50-server/persistence"
)

type handlerFunc = func(w http.ResponseWriter, req *http.Request)

func AddEntry(db *persistence.DB) handlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var death model.PlayerDeath
		err := json.NewDecoder(req.Body).Decode(&death)
		if err != nil {
			fmt.Printf("couldn't parse message\n")
		}
		fmt.Fprintf(w, "done")
		RecordPlayerDeath(death, db)
	}
}

func GetEntrySummary(db *persistence.DB) handlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		shrines := FetchAllShrines(db)
		err := json.NewEncoder(w).Encode(shrines)
		if err != nil {
			fmt.Printf("failed to write message\n")
			fmt.Fprintf(w, "failed")
		}
	}
}

func RunServer(addr string, db *persistence.DB) error {
	http.HandleFunc("/record_death", AddEntry(db))
	http.HandleFunc("/get_shrines", GetEntrySummary(db))
	return http.ListenAndServe(addr, nil)
}
