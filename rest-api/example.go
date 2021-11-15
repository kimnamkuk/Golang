package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var users = map[string]*User{}

type User struct {
	First_name string
	Last_name  string
}

type testHandler struct{}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Add("content-type", "application/json")
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Bad Requests:  ", err)
			return
		}
	case http.MethodPost:
		user := new(User)
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Bad Requests:  ", err)
			return
		}
		data, _ := json.Marshal(user)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, string(data))

		users[user.First_name] = user
	}

}

func test() {
	mux := http.NewServeMux()
	mux.Handle("/Users", &testHandler{})

	http.ListenAndServe(":8080", mux)
}
