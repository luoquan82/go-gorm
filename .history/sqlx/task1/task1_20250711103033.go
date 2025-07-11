package task1

import "github.com/jmoiron/sqlx"

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func initEmployee(d *sqlx.DB) {
}
