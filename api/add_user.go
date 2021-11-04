package api

import (
  "fmt"

  "github.com/go-telegram-bot-api/telegram-bot-api"

  "genshin_help_bot/GenshinDB"
)

func (uDB UserDB)add_user(update tgbotapi.Update) {
  user := update.Message.From
  values := fmt.Sprintf("(%d,'%s','%s')", user.ID, user.FirstName, user.LastName)
  GenshinDB.Insert(uDB.DB, uDB.Table, "ID,First_name,Last_name", values)
}