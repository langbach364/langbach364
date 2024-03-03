package main

import (
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

type account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type quantity struct {
	Start int `json:"start"`
	Limit int `json:"limit"`
}

func sign_Up(jsonData []byte) string {
	var Account account
	err := json.Unmarshal(jsonData, &Account)
	check_err(err)

	if check_email(Account.Email) {
		Account.Email = split_string(Account.Email)
	}
	hashedPassword := encode_data(Account.Email, Account.Password, 2)

	db, err := sql.Open("mysql", "root:@ztegc4DF9F4E@tcp(localhost)/store")
	check_err(err)

	_, err = db.Exec("INSERT INTO account(username, password) VALUES(?, ?)", Account.Email, hashedPassword)
	check_err(err)
	defer db.Close()

	check := true
	JsonData,err := json.Marshal(check)
	check_err(err)


	return string(JsonData)
}

func sign_In(jsonData []byte) string {
	var Account account
	err := json.Unmarshal(jsonData, &Account)
	check_err(err)

	hashPassword := encode_data(Account.Password, Account.Password, 2)

	check := check_login(Account.Email, hashPassword)
	JsonData, err := json.Marshal(check)
	check_err(err)

	return string(JsonData)
}

func handle_search(jsonData []byte) string {
	var data string
	err := json.Unmarshal(jsonData, &data)
	check_err(err)

	list := search_data(data)

	JsonData, err := json.Marshal(list)
	check_err(err)

	return string(JsonData)
}

func view_more(jsonData []byte) string {
	var Quantidy quantity
	err := json.Unmarshal(jsonData, &Quantidy)
	check_err(err)

	data := get_data(Quantidy.Start, Quantidy.Limit)

	JsonData, err := json.Marshal(data)
	check_err(err)

	return string(JsonData)
}
