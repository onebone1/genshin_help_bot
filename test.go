package main

import (
	//"fmt"

	// "github.com/joho/godotenv"

	"genshin_help_bot/cron"
	"genshin_help_bot/api"
	"genshin_help_bot/bot_func"
	// "genshin_help_bot/account"
)

func main() {
	_, updates := bot_func.Bot_init()
	users := api.Init()
	defer users.DB.Close()
	go cron.Crontab(users.User_signin, 0, 0)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.From.ID != 343140476 {
			continue
		}
		go users.Main_API(update)
	}
	// rows := GenshinDB.FindUser(users.DB, "ID", "")
	// for rows.Next() {
	//   var id string
	//   _ = rows.Scan(&id)
	//   fmt.Println(id)
	// }
	/*
	  var users [][]string
	  cols := "ID,First_name,State"
	  rows := GenshinDB.FindUser(DB, cols, "First_name='3rd'")
	  for rows.Next() {
	    var (
	      id string
	      fn string
	      s string
	    )
	    _ = rows.Scan(&id, &fn, &s)
	    fmt.Println(id, fn, s)
	    var user_info []string
	    user_info = append(user_info, id)
	    user_info = append(user_info, fn)
	    user_info = append(user_info, s)
	    users = append(users, user_info)
	  }
	  fmt.Println(users)
	  fmt.Println(len(users))
	  defer rows.Close()
	  /*
	  var user account.User
	  user.ID  = 123
	  user.FirstName = "test user 1"
	  user.LastName = "test last name"
	  user.Uid = "800800800"
	  user.Account_id = "987654321"
	  user.Cookie_token = "some token"
	  user.State = 1.0
	  GenshinDB.AddUser(DB, user)
	*/
}
