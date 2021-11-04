package api

import (
  "fmt"
  "log"

  "genshin_help_bot/GenshinDB"
)

func (uDB UserDB)if_exist(ID int)(b bool) {
  condition := fmt.Sprintf("ID=%d", ID)
  rows := GenshinDB.Select(uDB.DB, uDB.Table, "ID", condition)
  count := 0
  for rows.Next() {
    count += 1
  }
  if count == 1 {
    return true
  }else if count == 0{
    log.Println("User not found!")
  }else {
    log.Println("Multiple user found!")
  }
  return false
}

func (uDB UserDB)get_state(ID int)(state float64) {
  condition := fmt.Sprintf("ID=%d", ID)
  rows := GenshinDB.Select(uDB.DB, uDB.Table, "State", condition)
  for rows.Next() {
    _ = rows.Scan(&state)
  }
  return state
}

func (uDB UserDB)get_game_info(ID int)(UID string, Account_id string, Cookie_token string) {
  condition := fmt.Sprintf("ID=%d", ID)
  rows := GenshinDB.Select(uDB.DB, uDB.Table, "Uid,Account_id,Cookie_token", condition)
  for rows.Next() {
    _ = rows.Scan(&UID, &Account_id, &Cookie_token)
  }
  return UID, Account_id, Cookie_token
}