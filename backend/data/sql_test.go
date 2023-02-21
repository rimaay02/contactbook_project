package data

import (
	"context"
	"fmt"
	"testing"
)

func TestExcecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	first_name := "henry"
	last_name := "doe"
	phone_number := "012233"
	email := "henry.doe@example.com"

	script := "INSERT INTO contact (first_name, last_name, phone_number, email) VALUES(?,?,?,?)"

	_, err := db.ExecContext(ctx, script, first_name, last_name, phone_number, email)
	if err != nil {
		panic(err)
	}

	fmt.Println("input success")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, first_name FROM contact"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, first_name string
		err := rows.Scan(&id, &first_name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("First Name:", first_name)
	}
}