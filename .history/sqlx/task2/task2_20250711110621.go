package task2

import (
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func initBooks(db *sqlx.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		author VARCHAR(20) NOT NULL,
		price DECIMAL(10, 2) NOT NULL
	);
	`
	_, err := db.Exec(schema)
	if err != nil {
		panic(err)
	}

	books := []Book{
		{Title: "Go语言编程指南", Author: "马克·贝茨"},
		{Title: "Go并发编程"},
	}
}

func SelectBookOver50(d *sqlx.DB) {
	initBooks(d)
	books := []Book{}
	d.Select(&books, "select * from books where price>?", 50)
}
