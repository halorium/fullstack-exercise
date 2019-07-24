package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

var tests = []struct {
  name     string
  query    string
  response string
}{
  {
    name:     "all people",
    query:    "",
    response: `[{"id":1,"name":"Vinny Balaison","imgUrl":"https://robohash.org/avoluptatesut.jpg?size=64x64\u0026set=set1","location":"Batanamang","colors":null},{"id":2,"name":"Rourke Luckin","imgUrl":"https://robohash.org/utquamreiciendis.jpg?size=64x64\u0026set=set1","location":"Iguig","colors":null},{"id":3,"name":"Brianna McElree","imgUrl":"https://robohash.org/omnisbeataeest.jpg?size=64x64\u0026set=set1","location":"Formiga","colors":null}]`,
  },
  {
    name:     "subset",
    query:    "bal",
    response: `[{"id":1,"name":"Vinny Balaison","imgUrl":"https://robohash.org/avoluptatesut.jpg?size=64x64\u0026set=set1","location":"Batanamang","colors":null}]`,
  },
}

type mockDB struct{}

func (mdb *mockDB) GetPeople(query string) ([]*Person, error) {
  people := make([]*Person, 0)
  if query == "" {
    people = append(people,
      &Person{
        ID:       1,
        Name:     "Vinny Balaison",
        IMGURL:   "https://robohash.org/avoluptatesut.jpg?size=64x64&set=set1",
        Location: "Batanamang",
      },
      &Person{
        ID:       2,
        Name:     "Rourke Luckin",
        IMGURL:   "https://robohash.org/utquamreiciendis.jpg?size=64x64&set=set1",
        Location: "Iguig",
      },
      &Person{
        ID:       3,
        Name:     "Brianna McElree",
        IMGURL:   "https://robohash.org/omnisbeataeest.jpg?size=64x64&set=set1",
        Location: "Formiga",
      },
    )
  } else {
    people = append(people,
      &Person{
        ID:       1,
        Name:     "Vinny Balaison",
        IMGURL:   "https://robohash.org/avoluptatesut.jpg?size=64x64&set=set1",
        Location: "Batanamang",
      },
    )
  }

  return people, nil
}

func TestGetPeople(t *testing.T) {
  for _, test := range tests {
    t.Run(test.name, func(t *testing.T) {
      rec := httptest.NewRecorder()
      path := "/people"
      if test.query != "" {
        path = path + "?q=" + test.query
      }
      req, err := http.NewRequest("GET", path, nil)
      if err != nil {
        t.Fatal(err)
      }

      env := &Env{&mockDB{}}
      handler := http.HandlerFunc(GetPeople(env))

      handler.ServeHTTP(rec, req)

      status := rec.Code
      if status != http.StatusOK {
        t.Errorf("unexpected status code: expected %v got %v", http.StatusOK, status)
      }

      if rec.Body.String() != test.response {
        t.Errorf("unexpected body: expected:\n%v\ngot:\n%v\n", test.response, rec.Body.String())
      }
    })
  }
}
