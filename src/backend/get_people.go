package main

import (
  "encoding/json"
  "net/http"
)

func GetPeople(env *Env) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")

    query := r.URL.Query().Get("q")

    people, err := env.db.GetPeople(query)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    serialized, err := json.Marshal(people)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(serialized)
  }
}
