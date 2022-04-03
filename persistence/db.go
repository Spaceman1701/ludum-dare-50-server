package persistence

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Spaceman1701/ludum-dare-50-server/model"
)

func NewDb() (*gorm.DB, error) {
	fmt.Println("starting ld50 server")
	dsn := "host=localhost user=postgres password=foobar dbname=ld50 port=5432 sslmode=disable TimeZone=America/Denver"
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
