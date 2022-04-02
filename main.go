package main

import (
	"fmt"

	"github.com/Spaceman1701/ludum-dare-50-server/persistence"
	"github.com/Spaceman1701/ludum-dare-50-server/web"
)

func main() {
	db, err := persistence.NewDb()
	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	fmt.Println("successfully connected to db, starting server")

	if err = web.RunServer(":8090"); err != nil {
		panic(err)
	}
}
