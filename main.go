package main

import (
	"time"

	"genshin_help_bot/account"
	"genshin_help_bot/bot_func"
)

func Signin_crontab() {
	for true {
		t := time.Now()
		loc, _ := time.LoadLocation("Asia/Taipei")
		n := t.In(loc)
		hour := n.Hour()
		min := n.Minute()
		if hour == 0 && min == 0 {
			account.Accs.Signin()
			time.Sleep(23 * time.Hour)
		}
	}
}

func main() {
	account.Accs.Init()
	_, updates := bot_func.Bot_init()
	go Signin_crontab()

	for update := range updates {
		if update.Message != nil && !update.Message.From.IsBot {
			account.Accs.Acc_main(update)
		}
	}
}
