package model

import (
	"time"
)

const WorldWidth = 256
const WorldHeight = 256

const ShrineAreaOfEffect = 10

type WorldPosVector struct {
	X int
	Y int
}

type WorldPos = int

func WorldPosToVector(w WorldPos) WorldPosVector {
	var vec WorldPosVector
	vec.X = int(w) % WorldWidth
	vec.Y = int(w) / WorldWidth
	return vec
}

type PlayerDeath struct {
	Pos         WorldPos
	Usrename    string
	Time        time.Time
	Sacrifice   bool
	UsedShrines []uint
}

type ShrineState int

const (
	Potential ShrineState = iota
	Realized
)

type Player struct {
	Username string `gorm:"primaryKey"`
}

type Shrine struct {
	ID    uint `gorm:"primaryKey"`
	Pos   WorldPos
	Power int
	State ShrineState

	CreatedByUsername string
	CreatedBy         Player

	Contributors []Player `gorm:"many2many:contributors"`
}

type ShrineList struct {
	Shrines []Shrine
}
