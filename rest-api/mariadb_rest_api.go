package main

import (
	"encoding/json"
	"net/http"
)

type httpHandler struct{}

mariadb := new(MariaDB) 

func (h *httpHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Add("content-type", "application/json")
		err := json.NewEncoder(w).Encode()
	case http.MethodPost:
	}
}

func setMariaDB(m *MariaDB)
