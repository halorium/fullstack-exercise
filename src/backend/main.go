package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  db, err := NewDB()

  if err != nil {
    log.Panic(err)
  }

  env := &Env{db}

  router := mux.NewRouter()

  router.HandleFunc("/people", GetPeople(env)).Methods("GET", "OPTIONS")

  log.Fatal(
    http.ListenAndServe(":4000", router),
  )
}
