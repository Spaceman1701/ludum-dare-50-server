package web

import (
	"github.com/Spaceman1701/ludum-dare-50-server/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func getOrCreatePlayer(username string, db *gorm.DB) model.Player {
	var player model.Player
	res := db.First(&player, "Username = ?", username)
	if res.Error != nil {
		player = model.Player{
			Username: username,
		}
		db.Create(&player)
	}
	return player
}

func RecordPlayerDeath(death model.PlayerDeath, db *gorm.DB) {
	tx := db.Begin()
	defer tx.Commit()

	shrines := make([]model.Shrine, 0)
	player := getOrCreatePlayer(death.Username, tx)
	tx.Preload(clause.Associations).Find(&shrines)

	diedInShrine := false
	diedInBufferZone := false
	for _, s := range shrines {
		if s.IsPointInside(death.Pos) {
			contrib := model.ComputeShrineContribution(death, &s)
			s.Power += contrib
			if death.Sacrifice {
				s.Contributors = append(s.Contributors, player)
			}
			tx.Save(&s)
			diedInShrine = true
			break
		}
		if s.IsPointInBufferZone(death.Pos) {
			diedInBufferZone = true
		}
	}

	if !diedInShrine && !diedInBufferZone {
		shrine := model.CreateNewPotentialShrine(death, player)
		tx.Save(&shrine)
	}
}

func FetchAllShrines(db *gorm.DB) model.ShrineList {
	shrines := make([]model.Shrine, 0)
	db.Preload(clause.Associations).Find(&shrines)

	return model.ShrineList{Shrines: shrines}
}
