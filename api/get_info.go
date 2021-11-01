package api

import (
  "genshin_help_bot/GenshinDB"
)

func (uDB UserDB)get_state(ID int)(state float64) {
  rows := GenshinDB.Select(uDB.DB, uDB.Table, "State", "ID="+string(ID))
  for rows.Next() {
    _ = rows.Scan(&state)
  }
  return state
}

func (uDB UserDB)get_game_info(ID int)(UID string, Account_id string, Cookie_token string) {
  rows := GenshinDB.Select(uDB.DB, uDB.Table, "Uid,Account_id,Cookie_token", "ID="+string(ID))
  for rows.Next() {
    _ = rows.Scan(&UID, &Account_id, &Cookie_token)
  }
  return UID, Account_id, Cookie_token
}