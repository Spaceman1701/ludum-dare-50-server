package model

import "time"

const WorldWidth = 256
const WorldHeight = 256

type WorldPosVector struct {
	X int
	Y int
}

type WorldPos uint

func (w WorldPos) ToVector() WorldPosVector {
	var vec WorldPosVector
	vec.X = int(w) % WorldWidth
	vec.Y = int(w) / WorldWidth
	return vec
}

type TileVisit struct {
	ID   uint     `gorm:"primaryKey"`
	Pos  WorldPos //row major
	Time time.Time
}

type GameHistory struct {
	ID uint `gorm:"primaryKey"`
	//TODO
	TouchedPositions []TileVisit `gorm:"foreignKey:ID"`
	DeathLocation    WorldPos
}

type HealingShrine struct {
	ID       uint `gorm:"primaryKey"`
	Location WorldPos
	Power    int
}

type WearMap struct {
	Amounts []int
}

type GameState struct {
	Shrines []HealingShrine
	Wear    WearMap
}
