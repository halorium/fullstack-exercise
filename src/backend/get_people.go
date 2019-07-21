package main

import (
  "encoding/json"
  "net/http"
)

func getPeople(responseWriter http.ResponseWriter, request *http.Request) {
  responseWriter.Header().Set("Access-Control-Allow-Origin", "*")

  db, err := newDBConnection()

  if err != nil {
    http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
    return
  }

  defer db.Close()

  var people []*Person

  query := request.URL.Query().Get("q")

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

  serialized, err := json.Marshal(people)

  if err != nil {
    http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
    return
  }

  responseWriter.Write(serialized)
}
