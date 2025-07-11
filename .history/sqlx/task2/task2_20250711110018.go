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
		name VARCHAR(100) NOT NULL,
		department VARCHAR(100) NOT NULL,
		salary DECIMAL(10, 2) NOT NULL
	);
	`
	_, err := db.Exec(schema)
	if err != nil {
		panic(err)
	}
}

func SelectBookOver50(d *sqlx.DB) {
	initBooks(d)
	books := []Book{}
	d.Select(&books, "select * from books where price>?", 50)
}
