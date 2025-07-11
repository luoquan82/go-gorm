package task2

import "github.com/jmoiron/sqlx"

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func initBooks(db *sqlx.DB) {

}

func SelectBookOver50(d *sqlx.DB) {
	books := []Book{}
	d.Select(&books, "select * from books where price>?", 50)
}
