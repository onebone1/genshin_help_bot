package main

import (
	"time"

	"genshin/account"
	"genshin/bot_func"
)

func Signin_crontab() {
  for true {
    hour := time.Now().Hour()
    min := time.Now().Minute()
    if hour == 0 && min == 0 {
      account.Accs.Signin()
      time.Sleep(23*time.Hour)
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
