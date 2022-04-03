package web

import (
	"github.com/Spaceman1701/ludum-dare-50-server/model"
	"github.com/Spaceman1701/ludum-dare-50-server/persistence"
)

func RecordPlayerDeath(death model.PlayerDeath, db *persistence.DB) {

}

func FetchAllShrines(db *persistence.DB) model.ShrineList {
	return model.ShrineList{}
}
