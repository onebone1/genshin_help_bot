package api

import (
	"fmt"

	"genshin_help_bot/GenshinDB"
)

func (uDB UserDB) set_state(ID int, state float64) {
	str := fmt.Sprintf("State=%f", state)
	condition := fmt.Sprintf("ID=%d", ID)
	GenshinDB.Update(uDB.DB, uDB.Table, str, condition)
}
