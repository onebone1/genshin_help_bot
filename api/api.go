package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"

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
	DB    *sql.DB
	Table string
}

func Init() (uDB UserDB) {
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

func (uDB UserDB) if_exist(ID int) (b bool) {
	condition := fmt.Sprintf("ID=%d", ID)
	rows := GenshinDB.FindUser(uDB.DB, "ID", condition)
	count := 0
	for rows.Next() {
		count += 1
	}
	if count == 0 {
		log.Println("User not found!")
	} else if count > 1 {
		log.Println("Multiple user found!")
	} else {
		log.Println("User found!")
		return true
	}
	return false
}

func (uDB UserDB) Main_API(update tgbotapi.Update) {
	user := update.Message.From
	text := update.Message.Text
	//chatID := update.Message.Chat.ID
	if !uDB.if_exist(user.ID) {
		bot_func.TGBot.SendMessage(int64(user.ID), "First use")
		return
	}
	log.Println(uDB.Table)
	state := uDB.get_state(user.ID)
	if state == 1.0 {
		// main functions
		if text == "/devsignin" {
			uDB.User_signin()
		} else if text == "/gift" {
			bot_func.TGBot.SendMessage(int64(user.ID), "請輸入兌換碼")
			uDB.set_state(user.ID, 3.0)
		} else if text == "/broadcast" && user.ID == 343140476 {
			bot_func.TGBot.SendMessage(int64(user.ID), "開始廣播，輸入 /end 結束廣播")
			uDB.set_state(user.ID, 4.0)
		}
	} else if int(state) == 2 {

	} else if state == 3.0 {
		uDB.Gift(user.ID, text)
		// str := fmt.Sprintf("使用者 耕銘 已經幫您使用以下兌換碼\n%s", text)
		// bot_func.TGBot.SendMessage(int64(user.ID), str)
		uDB.set_state(user.ID, 1.0)
	} else if state == 4.0 {
		if text == "/end" {
			bot_func.TGBot.SendMessage(int64(user.ID), "廣播結束")
			uDB.set_state(user.ID, 1.0)
			return
		}
		uDB.broadcast(text)
	}
}
