package persistence

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Spaceman1701/ludum-dare-50-server/model"
)

type DB interface {
	WriteEntry(model.GameHistory) error
	ReadAllEntries() ([]model.GameHistory, error)
}

type gormDb struct {
	db *gorm.DB
}

func (d *gormDb) WriteEntry(e model.GameHistory) error {
	return d.db.Create(&e).Error
}

func (d *gormDb) ReadAllEntries() ([]model.GameHistory, error) {
	history := make([]model.GameHistory, 0)
	res := d.db.Find(&history)
	return history, res.Error
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
