package main

import (
  "os"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Datastore interface {
  GetPeople(query string) ([]*Person, error)
}

type DB struct {
  *gorm.DB
}

func NewDB() (*DB, error) {
  username := os.Getenv("MYSQL_USER")
  password := os.Getenv("MYSQL_PASSWORD")
  dbname := os.Getenv("MYSQL_DATABASE")

  connectionString := username + ":" + password + "@tcp(db:3306)/" + dbname

  db, err := gorm.Open("mysql", connectionString)

  if err != nil {
    return nil, err
  }

  return &DB{db}, nil
}

func (db *DB) GetPeople(query string) ([]*Person, error) {
  var people []*Person

  if query != "" {
    query = "%" + query + "%"
    db.Preload("PeopleColors").Where("LOWER(name) LIKE ?", query).Find(&people)
  } else {
    db.Preload("PeopleColors").Find(&people)
  }

  for _, person := range people {
    for _, color := range person.PeopleColors {
      person.Colors = append(person.Colors, color.Color)
    }
  }

  return people, nil
}
