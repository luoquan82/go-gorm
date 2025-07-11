package task1

import "github.com/jmoiron/sqlx"

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func initEmployee(d *sqlx.DB) {
	// Create the employees table if it doesn't exist
	schema := `
	CREATE TABLE IF NOT EXISTS employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		department VARCHAR(100) NOT NULL,
		salary DECIMAL(10, 2) NOT NULL
	);`
	_, err := d.Exec(schema)
	if err != nil {
		panic(err)
	}

	// Insert sample data into the employees table
	employees := []Employee{
		{Name: "Alice", Department: "Engineering", Salary: 75000.00},
		{Name: "Bob", Department: "Marketing", Salary: 60000.00},
		{Name: "Charlie", Department: "Sales", Salary: 55000.00},
	}

	for _, emp := range employees {
		_, err = d.NamedExec("INSERT INTO employees (name, department, salary) VALUES (:name, :department, :salary)", &emp)
		if err != nil {
			panic(err)
		}
	}
}

func Run(db *sqlx.DB) {
	initEmployee(db)
}
