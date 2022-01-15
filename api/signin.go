package api

import (
	"fmt"
	"genshin_help_bot/GenshinDB"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func signin(Account_id string, Cookie_token string) {
	URL := "https://hk4e-api-os.mihoyo.com/event/sol/sign?lang=zh-tw"
	data := strings.NewReader(`{"act_id":"e202102251931481"}`)
	req, err := http.NewRequest("POST", URL, data)
	if err != nil {
		fmt.Println(err)
	}

	req.AddCookie(&http.Cookie{Name: "account_id", Value: Account_id})
	req.AddCookie(&http.Cookie{Name: "cookie_token", Value: Cookie_token})
	log.Println("\n", req, "\n")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	// return string(body)
}

func (uDB UserDB) User_signin() {
	rows := GenshinDB.FindUser(uDB.DB, "Account_id,Cookie_token", "")
	var (
		Account_id   string
		Cookie_token string
	)
	for rows.Next() {
		_ = rows.Scan(&Account_id, &Cookie_token)
		signin(Account_id, Cookie_token)
	}
}
