package model

const ShrineAreaOfEffect = 10
const ShrineBufferArea = 20

type Vec2 struct {
	X int
	Y int
}

type PlayerDeath struct {
	Pos         Vec2
	Username    string
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
	Pos   Vec2 `gorm:"embedded"`
	Power int
	State ShrineState

	CreatedByUsername string
	CreatedBy         Player

	Contributors []Player `gorm:"many2many:contributors"`
}

func dist2(a Vec2, b Vec2) int {
	xDist := a.X - b.X
	yDist := a.Y - b.Y
	return xDist*xDist + yDist*yDist
}

func (s *Shrine) IsPointInBufferZone(pos Vec2) bool {
	buffer := GetConfig().Shrine.Buffer
	return dist2(s.Pos, pos) < (buffer * buffer)
}

func (s *Shrine) IsPointInside(pos Vec2) bool {
	aoe := GetConfig().Shrine.Radius
	return dist2(s.Pos, pos) < (aoe * aoe)
}

type ShrineList struct {
	Shrines []Shrine
}

func ComputeShrineContribution(death PlayerDeath, shrine *Shrine) int {
	if death.Sacrifice {
		return GetConfig().Shrine.BaseSacrificeContrib
	} else {
		return GetConfig().Shrine.BaseDeathContrib
	}
}

func CreateNewPotentialShrine(death PlayerDeath, player Player) Shrine {
	return Shrine{
		Pos:       death.Pos,
		Power:     GetConfig().Shrine.SpawnPower,
		State:     Potential,
		CreatedBy: player,
	}
}

func ComputeShrineCost(death PlayerDeath, shrine *Shrine) int {
	return GetConfig().Shrine.UsageCost
}
