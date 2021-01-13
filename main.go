package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	dbCon = initDb()
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")

	http.ListenAndServe(":5055", r)
}

func initDb() *sql.DB {
	db, e := sql.Open("mysql", "test:PassworD12312312?@tcp(127.0.0.1)/shopping")
	if e != nil {
		panic(e.Error())
	}
	return db
}
