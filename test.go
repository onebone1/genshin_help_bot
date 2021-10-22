package main

import (
	"fmt"

	// "github.com/joho/godotenv"

  "genshin_help_bot/GenshinDB"
  // "genshin_help_bot/account"
)

func main() {
	DB := GenshinDB.Init()
  defer DB.Close()
  cols := "ID,First_name"
  rows := GenshinDB.FindUser(DB, cols, "ID=343140476")
  for rows.Next() {
    var (
      id int
      fn string
    )
    _ = rows.Scan(&id, &fn)
    fmt.Println(id, fn)
  }
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
  GenshinDB.DeleteUser(DB, "user", "ID=456")
}
