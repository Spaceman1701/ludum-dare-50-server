package web

import (
	"fmt"

	"github.com/Spaceman1701/ludum-dare-50-server/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type interactionResult struct {
	shrineModified     bool
	diedInBufferZone   bool
	shouldDeleteShrine bool
}

func handleShrineInteraction(death model.PlayerDeath, shrine *model.Shrine, player model.Player) interactionResult {

	diedInShrine := shrine.IsPointInside(death.Pos)
	diedInBuffer := shrine.IsPointInBufferZone(death.Pos)
	usedShrine := playerUsedShrine(death, shrine)

	res := interactionResult{
		diedInBufferZone:   diedInBuffer,
		shrineModified:     false,
		shouldDeleteShrine: false,
	}

	if diedInShrine {
		shrine.Power += model.ComputeShrineContribution(death, shrine)
		if death.Sacrifice {
			shrine.Contributors = append(shrine.Contributors, player)
		}

		res.shrineModified = true
	}

	if usedShrine {
		shrine.Power -= model.ComputeShrineCost(death, shrine)

		res.shrineModified = true
	}

	if !diedInShrine && shrine.State == model.Potential {
		//potential shrines decay if not contributed to
		shrine.Power -= model.GetConfig().Shrine.DeclineRate

		res.shrineModified = true
	}

	if res.shrineModified {
		if shrine.Power < model.GetConfig().Shrine.DestroyThreshold {
			res.shouldDeleteShrine = true
		} else if shrine.State == model.Potential && shrine.Power >= model.GetConfig().Shrine.UpgradeThreshold {
			shrine.State = model.Realized
		}
	}
	return res
}

func handleStateUpdate(db *gorm.DB, death model.PlayerDeath) {
	tx := db.Begin()
	defer func() {
		tx.Commit()
		fmt.Println("committed update to database")
	}()

	shrines := make([]model.Shrine, 0)
	player := getOrCreatePlayer(death.Username, tx)
	tx.Preload(clause.Associations).Find(&shrines)

	createPotential := true
	for _, shrine := range shrines {
		res := handleShrineInteraction(death, &shrine, player)
		if res.shouldDeleteShrine {
			tx.Delete(&shrine)
		} else if res.shrineModified {
			tx.Save(&shrine)
		}
		if res.diedInBufferZone {
			createPotential = false
		}
	}

	if createPotential {
		s := model.CreateNewPotentialShrine(death, player)
		tx.Save(&s)
	}
}

func AsyncShrineUpdate(db *gorm.DB, eventQueue chan model.PlayerDeath) {
	for {
		death := <-eventQueue
		handleStateUpdate(db, death)
	}
}
