package model

import (
	"time"
)

const WorldWidth = 256
const WorldHeight = 256

const ShrineAreaOfEffect = 10
const ShrineBufferArea = 20

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
	Username    string
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

func dist2(a WorldPosVector, b WorldPosVector) int {
	xDist := a.X - b.X
	yDist := a.Y - b.Y
	return xDist*xDist + yDist*yDist
}

func (s *Shrine) IsPointInBufferZone(pos WorldPos) bool {
	sVec := WorldPosToVector(s.Pos)
	pVec := WorldPosToVector(pos)
	return dist2(sVec, pVec) < (ShrineBufferArea * ShrineBufferArea)
}

func (s *Shrine) IsPointInside(pos WorldPos) bool {
	sVec := WorldPosToVector(s.Pos)
	pVec := WorldPosToVector(pos)

	return dist2(sVec, pVec) < (ShrineAreaOfEffect * ShrineAreaOfEffect)
}

type ShrineList struct {
	Shrines []Shrine
}

func ComputeShrineContribution(death PlayerDeath, shrine *Shrine) int {
	if death.Sacrifice {
		return 100
	} else {
		return 10
	}
}

func CreateNewPotentialShrine(death PlayerDeath, player Player) Shrine {
	return Shrine{
		Pos:       death.Pos,
		Power:     10,
		State:     Potential,
		CreatedBy: player,
	}
}

func ComputeShrineCost(death PlayerDeath, shrine *Shrine) int {
	return 10
}
