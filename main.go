package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("starting ld50 server")
	dsn := "host=localhost user=postgres password=foobar dbname=ld50 port=5432 sslmode=disable TimeZone=America/Denver"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
