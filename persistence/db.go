package persistence

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Spaceman1701/ludum-dare-50-server/model"
)

func NewDb() (*gorm.DB, error) {
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}

	fmt.Println("starting ld50 server")
	dsn := fmt.Sprintf("host=%s user=postgres password=%s dbname=ld50 port=5432 sslmode=disable TimeZone=America/Denver", host, password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Player{}, &model.Shrine{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
