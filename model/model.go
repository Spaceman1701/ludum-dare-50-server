package model

import "time"

type TileVisit struct {
	ID   uint `gorm:"primaryKey"`
	Pos  uint
	Time time.Time
}

type GameHistory struct {
	ID uint `gorm:"primaryKey"`
	//TODO
	TouchedPositions TileVisit `gorm:"foreignKey:ID"`
}

type GameHistorySummary struct {
}
