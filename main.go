package main

import (
	"fmt"

	"github.com/Spaceman1701/ludum-dare-50-server/persistence"
)

func main() {
	db, err := persistence.NewDb()
	if err != nil {
		panic(err)
	}

	fmt.Println(db)
}
