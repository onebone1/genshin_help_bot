package main

import (
	"github.com/roylee0704/gron"
	"github.com/roylee0704/gron/xtime"

	"genshin/account"
	"genshin/bot_func"
)

func main() {
	account.Accs.Init()
	_, updates := bot_func.Bot_init()
	cron := gron.New()
	cron.AddFunc(gron.Every(1*xtime.Day).At("00:00"), account.Accs.Signin)
	cron.Start()

	for update := range updates {
		if update.Message != nil && !update.Message.From.IsBot {
			account.Accs.Acc_main(update)
		}
	}
}
