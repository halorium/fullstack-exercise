package main

import (
  "os"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func newDBConnection() (*gorm.DB, error) {
  username := os.Getenv("MYSQL_USER")
  password := os.Getenv("MYSQL_PASSWORD")
  dbname := os.Getenv("MYSQL_DATABASE")

  connectionString := username + ":" + password + "@tcp(db:3306)/" + dbname

  db, err := gorm.Open("mysql", connectionString)

  if err != nil {
    return nil, err
  }

  return db, nil
}
