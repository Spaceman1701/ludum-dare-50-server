package model

import "time"

type TilePos struct {
	ID uint
	X  int `gorm:"primaryKey;autoIncrement:false"`
	Y  int `gorm:"primaryKey:autoIncrement:false"`
}

type TileVisit struct {
	ID   uint    `gorm:"primaryKey"`
	Pos  TilePos `gorm:"foreignKey:ID"`
	Time time.Time
}

type GameHistory struct {
	ID uint `gorm:"primaryKey"`
	//TODO
	TouchedPositions TileVisit `gorm:"foreignKey:ID"`
}

type GameHistorySummary struct {
}
