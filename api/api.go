package api

import (
  "fmt"
  "os"
  "log"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

  "github.com/go-telegram-bot-api/telegram-bot-api"

  "genshin_help_bot/GenshinDB"
  "genshin_help_bot/bot_func"
)

const (
  my_info     = 1
  change_info = 2
  gift        = 3
  devsignin   = 4
  devstate    = 5
)

type UserDB struct {
  DB *sql.DB
  Table string
}

func Init()(uDB UserDB) {
  uDB.DB = GenshinDB.Init()
  uDB.Table = os.Getenv("User_table")
  return uDB
}

func Instruction(text string) int {
  if text == "/my_info" || text == "/my_info"+bot_func.Bot_info.Username {
    return my_info
  } else if text == "/change_info" || text == "/change_info"+bot_func.Bot_info.Username {
    return change_info
  } else if text == "/gift" || text == "/gift"+bot_func.Bot_info.Username {
    return gift
  } else if text == "/devsignin" {
    return devsignin
  } else if text == "/devstate" {
    return devstate
  }
  return 0
}

func (uDB UserDB)if_exist(ID int)(b bool) {
  condition := fmt.Sprintf("ID=%d", ID)
  rows := GenshinDB.FindUser(uDB.DB, "ID", condition)
  count := 0
  for rows.Next() {
    count += 1
  }
  if count == 0 {
    log.Println("User not found!")
  }else if count > 1 {
    log.Println("Multiple user found!")
  }else {
    return true
  }
  return false
}

/*
func (uDB UserDB)get_state(ID int)(state float64) {
  rows := GenshinDB.FindUser(uDB.DB, "State", "ID="+string(ID))
  for rows.Next() {
    _ = rows.Scan(&state)
  }
  return state
}
*/

func (uDB UserDB)Main_API(update tgbotapi.Update) {
  user := update.Message.From
  //text := update.Message.Text
  //chatID := update.Message.Chat.ID
  if !uDB.if_exist(user.ID) {
    bot_func.TGBot.SendMessage(int64(user.ID), "First use")
    return
  }
  state := uDB.get_state(user.ID)
  if state == 1.0 {
    // main functions
  }
}