package main

import (
	"fmt"

	"github.com/Spaceman1701/ludum-dare-50-server/model"
	"github.com/Spaceman1701/ludum-dare-50-server/persistence"
	"github.com/Spaceman1701/ludum-dare-50-server/web"
)

func main() {
	db, err := persistence.NewDb()
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully connected to db, starting server")

	shrineUpdates := make(chan model.PlayerDeath)
	go web.AsyncShrineUpdate(db, shrineUpdates)

	if err = web.RunServer(":8090", db, shrineUpdates); err != nil {
		panic(err)
	}
}
