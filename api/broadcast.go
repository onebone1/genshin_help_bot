package api

import (
	// "log"

	"genshin_help_bot/GenshinDB"
	"genshin_help_bot/bot_func"
)

func (uDB UserDB) broadcast(text string) {
	rows := GenshinDB.FindUser(uDB.DB, "ID", "")
	var ID int64
	for rows.Next() {
		_ = rows.Scan(&ID)
		bot_func.TGBot.SendMessage(ID, text)
	}
}
