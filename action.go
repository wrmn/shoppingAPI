package main

import "database/sql"

func selectProduct(id string, db *sql.DB) (product, error) {
	item := product{}
	rowProduct, e := dbCon.Query("SELECT * FROM produk WHERE id=?", id)
	for rowProduct.Next() {
		e = rowProduct.Scan(
			&item.ID,
			&item.Name,
			&item.Price,
		)
	}
	return item, e
}
