package task1

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

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
		{Name: "张三", Department: "技术部", Salary: 75000.00},
		{Name: "李四", Department: "市场", Salary: 60000.00},
		{Name: "王五", Department: "技术部", Salary: 55000.00},
		{Name: "赵六", Department: "销售", Salary: 55000.00},
		{Name: "郑七", Department: "技术部", Salary: 55000.00},
		{Name: "朱八", Department: "运营", Salary: 55000.00},
		{Name: "孙九", Department: "技术部", Salary: 55000.00},
		{Name: "周是", Department: "采购", Salary: 55000.00},
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

	employees := []Employee{}
	err := db.Select(&employees, "SELECT * FROM employees")
	if err != nil {
		panic("error occur")
	}

	fmt.Printf("技术部人员: %#v", employees)
}
