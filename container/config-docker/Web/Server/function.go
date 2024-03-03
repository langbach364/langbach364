package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Type string
	Name string
}

func check_err(err error) {
	if err != nil {
		println(err)
		log.Fatal(err)
	}
}

func check_email(email string) bool {
	return len(email) >= 11
}

func split_string(data string) string {
	result := data[:len(data)-10]
	return result
}

func check_login(Username string, Password string) bool {
	db, err := sql.Open("mysql", "root:@ztegc4DF9F4E@tcp(localhost)/store")
	check_err(err)
	defer db.Close()

	var storedPassword string
	err = db.QueryRow("SELECT password FROM account WHERE username = ?", Username).Scan(&storedPassword)

	switch {
	case err == sql.ErrNoRows:
		return false

	case err != nil:
		check_err(err)

	default:
		if Password != storedPassword {
			return false
		}
	}
	return true
}

func get_data(start int, limit int) []Product {
	var list []Product
	db, err := sql.Open("mysql", "root:@ztegc4DF9F4E@tcp(localhost)/store")
	check_err(err)
	defer db.Close()

	rows, err := db.Query("SELECT name, type FROM product LIMIT ? OFFSET ? ", limit, start)
	check_err(err)

	var sum_object int
	err = db.QueryRow("SELECT COUNT(*) FROM (SELECT * FROM product LIMIT ? OFFSET ?) as product", limit, start).Scan(&sum_object)
	check_err(err)

	for i := 0; i < sum_object; i++ {
		var p Product
		rows.Next()
		rows.Scan(&p.Name, &p.Type)
		list = append(list, p)
	}
	return list
}

func check_space_last(data rune) bool {
	return data == ' '
}

func get_size_name(data string) int {
	for i := len(data) - 1; i > 0; i-- {
		if !check_space_last(rune(data[i])) {
			return i
		}
	}
	return 0
}

func reverse_string(data string) string {
	var result string
	for i := len(data) - 1; i >= 0; i-- {
		result += string(data[i])
	}
	return result
}

func get_name_string(data string, SavePointer *int) string {
	Size := get_size_name(data)

	var result string
	for i := Size; i >= 0; i-- {
		if check_space_last(rune(data[i])) {
			*SavePointer = i
			break
		}
		result += string(data[i])
	}
	return reverse_string(result)
}

func get_type_string(data string, SavePointer int) string {
	var result string

	data = strings.ToLower(data)

	for i := 0; i < SavePointer; i++ {
		result += string(data[i])
	}
	return result
}

func search_data(data string) []Product {
	var SavePointer int
	var list []Product
	Name := get_name_string(data, &SavePointer)
	Type := get_type_string(data, SavePointer)

	db, err := sql.Open("mysql", "root:@ztegc4DF9F4E@tcp(localhost)/store")
	check_err(err)
	defer db.Close()

	rows, err := db.Query("SELECT ?, ? from product", Name, Type)
	check_err(err)

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Name, &p.Type)
		check_err(err)
		list = append(list, p)
	}
	return list
}
