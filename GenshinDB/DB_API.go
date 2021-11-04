package GenshinDB

import (
  "fmt"
  "log"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func Insert(DB *sql.DB, table string, cols string, values string) {
  str := fmt.Sprintf("INSERT INTO %s (%s) VALUES%s;", table, cols, values)
  log.Println(str, "\n")
  Insert, err := DB.Query(str)
  if err != nil {
    log.Println("Insert error:", err)
  }else {
    log.Println("Insert successfully!")
  }
  Insert.Close()
}

func Select(DB *sql.DB, table string, cols string, conditions string)(Rows *sql.Rows) {
  str := fmt.Sprintf("SELECT %s FROM %s", cols, table)
  if conditions != "" {
    str = str + fmt.Sprintf(" WHERE %s", conditions)
  }
  Select, err := DB.Query(str + ";")
  if err != nil {
    log.Println("Select error:", err)
  }
  return Select
}

func Update(DB *sql.DB, table string, key_values string, conditions string) {
  str := fmt.Sprintf("UPDATE %s SET %s WHERE %s;", table, key_values, conditions)
  Update, err := DB.Query(str)
  if err != nil {
    log.Println("Update error:", err)
  }else {
    log.Println("Update successfully!")
  }
  Update.Close()
}

func Delete(DB *sql.DB, table string, conditions string) {
  str := fmt.Sprintf("DELETE FROM %s WHERE %s", table, conditions)
  Delete, err := DB.Query(str)
  if err != nil {
    log.Println("Delete error:", err)
  }else {
    log.Println("Delete successfully!")
  }
  Delete.Close()
}