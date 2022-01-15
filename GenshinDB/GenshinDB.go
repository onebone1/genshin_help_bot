package GenshinDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	//"genshin_help_bot/account"
)

type Users struct {
	DB *sql.DB
}

func Init() (DB *sql.DB) {
	_ = godotenv.Load(".test.env")
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

func (users Users) AddUser(table string, update tgbotapi.Update) {
	user := update.Message.From
	values := fmt.Sprintf("(%d,'%s','%s')", user.ID, user.FirstName, user.LastName)
	Insert(users.DB, table, "(ID,First_name,Last_name)", values)
}

func FindUser(DB *sql.DB, cols string, conditions string) (Rows *sql.Rows) {
	str := fmt.Sprintf("SELECT %s FROM user", cols)
	if conditions != "" {
		str = str + fmt.Sprintf(" WHERE %s", conditions)
	}
	log.Println("Query", str)
	query, err := DB.Query(str + ";")
	if err != nil {
		log.Println("Select error:", err)
	}
	return query
}

func UpdateUser(DB *sql.DB, table string, key_values string, condition string) {
	str := fmt.Sprintf("UPDATE %s SET %s WHERE %s;", table, key_values, condition)
	update, err := DB.Query(str)
	if err != nil {
		log.Println("Update error:", err)
	} else {
		log.Println("Update successfully!")
	}
	update.Close()
}

func DeleteUser(DB *sql.DB, table string, condition string) {
	str := fmt.Sprintf("DELETE FROM %s WHERE %s", table, condition)
	del, err := DB.Query(str)
	if err != nil {
		log.Println("Delete error", err)
	} else {
		log.Println("Delete successfully!")
	}
	del.Close()
}
