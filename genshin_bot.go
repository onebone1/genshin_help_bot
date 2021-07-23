package main

import (
	"time"

	"github.com/roylee0704/gron"

	"genshin/account"
	"genshin/bot_func"
)

func main() {
	account.Accs.Init()
	_, updates := bot_func.Bot_init()
	cron := gron.New()
	cron.AddFunc(gron.Every(1*time.Hour), account.Accs.Signin)
	cron.Start()

	for update := range updates {
		if update.Message != nil && !update.Message.From.IsBot {
			account.Accs.Acc_main(update)
		}
	}
}
