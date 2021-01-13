package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	product, err := selectProduct(id, dbCon)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(product)
}
