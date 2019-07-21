package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/people", getPeople).Methods("GET", "OPTIONS")

  log.Fatal(
    http.ListenAndServe(":4000", router),
  )
}
