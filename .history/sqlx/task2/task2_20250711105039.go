package task2

type Book struct {
	ID     int `db:"id"`
	Title  string
	Author string
	Price  float64
}
