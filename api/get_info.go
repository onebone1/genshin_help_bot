package api

import (
	"fmt"

	"genshin_help_bot/GenshinDB"
)

func (uDB UserDB) get_name(ID int) (name string) {
	condition := fmt.Sprintf("ID=%d", ID)
	rows := GenshinDB.Select(uDB.DB, uDB.Table, "First_name", condition)
	for rows.Next() {
		_ = rows.Scan(&name)
	}
	return name
}

func (uDB UserDB) get_state(ID int) (state float64) {
	condition := fmt.Sprintf("ID=%d", ID)
	rows := GenshinDB.Select(uDB.DB, uDB.Table, "State", condition)
	for rows.Next() {
		_ = rows.Scan(&state)
	}
	return state
}

func (uDB UserDB) get_game_info(ID int) (UID string, Account_id string, Cookie_token string) {
	condition := fmt.Sprintf("ID=%d", ID)
	rows := GenshinDB.Select(uDB.DB, uDB.Table, "Uid,Account_id,Cookie_token", condition)
	for rows.Next() {
		_ = rows.Scan(&UID, &Account_id, &Cookie_token)
	}
	return UID, Account_id, Cookie_token
}
