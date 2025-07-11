package task2

import "github.com/jmoiron/sqlx"

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func Run(d *sqlx.DB) {

}
