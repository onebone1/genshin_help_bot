package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"genshin_help_bot/bot_func"
	"genshin_help_bot/GenshinDB"
)

func Redeem(ID int, text string, UID string, Account_id string, Cookie_token string) {
	// UID, Account_id, Cookie_token := uDB.get_game_info(ID)
	uid := "uid=" + UID
	code := uid + "&cdkey=" + text
	url := "https://hk4e-api-os.mihoyo.com/common/apicdkey/api/webExchangeCdkey?region=os_asia&lang=zh-tw&game_biz=hk4e_global&" + code
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	req.AddCookie(&http.Cookie{Name: "account_id", Value: Account_id})
	req.AddCookie(&http.Cookie{Name: "cookie_token", Value: Cookie_token})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	bot_func.TGBot.SendMessage(int64(ID), string(body))
}

func (uDB UserDB) Gift(ID int, text string) {
	texts := strings.Split(text, "\n")
	username := uDB.get_name(ID)
	for _, str := range texts {
		UID, Account_id, Cookie_token := uDB.get_game_info(ID)
		Redeem(ID, str, UID, Account_id, Cookie_token)

		condition := fmt.Sprintf("ID != %d", ID)
		rows := GenshinDB.FindUser(uDB.DB, "ID, UID, Account_id, Cookie_token", condition)
		for rows.Next() {
			_ = rows.Scan(&ID, &UID, &Account_id, &Cookie_token)
			Redeem(ID, str, UID, Account_id, Cookie_token)
		}
		time.Sleep(5 * time.Second)
	}
	uDB.broadcast("使用者 " + username + " 已經幫您使用下列兌換碼\n" + text)
}