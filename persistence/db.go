package persistence

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Spaceman1701/ludum-dare-50-server/model"
)

type DB interface {
	WriteEntry(model.GameHistory) error
	ReadAllEntries() (model.GameHistorySummary, error)
}

type gormDb struct {
	db *gorm.DB
}

func (d *gormDb) WriteEntry(e model.GameHistory) error {
	return nil
}

func (d *gormDb) ReadAllEntries() (model.GameHistorySummary, error) {
	return model.GameHistorySummary{}, nil
}

func NewDb() (DB, error) {
	fmt.Println("starting ld50 server")
	dsn := "host=localhost user=postgres password=foobar dbname=ld50 port=5432 sslmode=disable TimeZone=America/Denver"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.GameHistory{})
	if err != nil {
		return nil, err
	}

	return &gormDb{
		db: db,
	}, nil
}
