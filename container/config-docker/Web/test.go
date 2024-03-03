package main

import (
	"database/sql"
	"fmt"
	"log"

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

func main() {
	List := get_data(2, 5)

	for _,x := range List {
		fmt.Println(x.Name, x.Type)
	}
}
