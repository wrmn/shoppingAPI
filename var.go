package main

import "database/sql"

var dbCon *sql.DB

type product struct {
	ID    int    `json:"ID"`
	Name  string `json:"nama"`
	Price string `json:"harga"`
}
