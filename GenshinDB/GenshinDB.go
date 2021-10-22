package GenshinDB

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"

	"genshin_help_bot/account"
)

func Init()(DB *sql.DB) {
	_ = godotenv.Load()
	UserName := os.Getenv("DB_user")
	Password := os.Getenv("DB_pass")
	Host := os.Getenv("DB_host")
	Port := os.Getenv("DB_port")
	DB_name := os.Getenv("DB_DB")
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", UserName, Password, Host, Port, DB_name)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
	return DB
}

func AddUser(DB *sql.DB, user account.User) {
	ID := user.ID
	FirstName := user.FirstName
	LastName := user.LastName
	UID := user.Uid
	Account_id := user.Account_id
	Cookie_token := user.Cookie_token
	State := user.State

	str := fmt.Sprintf("INSERT INTO user VALUES(%d,'%s','%s',%s,%s,'%s',%f)", ID, FirstName, LastName, UID, Account_id, Cookie_token, State)
	insert, err := DB.Query(str)
	if err != nil {
		log.Println("Insert error:", err)
	}
	insert.Close()
}

func FindUser(DB *sql.DB, cols string, conditions string)(Rows *sql.Rows) {
	str := fmt.Sprintf("SELECT %s FROM user", cols)
	if conditions != "" {
		str = str + fmt.Sprintf(" WHERE %s", conditions)
	}
	query, err := DB.Query(str + ";")
	if err != nil {
		log.Println("Select error:", err)
	}
  return query
}

func UpdateUser(DB *sql.DB, table string, col string, value string, condition string) {
  str := fmt.Sprintf("UPDATE %s SET %s=%s WHERE %s;", table, col, value, condition)
  update, err := DB.Query(str)
  if err != nil {
    log.Println("Update error:", err)
  }else {
    log.Println("Update successfully!")
  }
  update.Close()
}

func DeleteUser(DB *sql.DB, table string, condition string) {
  str := fmt.Sprintf("DELETE FROM %s WHERE %s", table, condition)
  del, err := DB.Query(str)
  if err != nil {
    log.Println("Delete error", err)
  }else {
    log.Println("Delete successfully!")
  }
  del.Close()
}