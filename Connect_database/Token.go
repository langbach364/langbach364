package main

import (
	"strconv"
	"database/sql"
)


func convert_int_to_string(data int) string {
	return strconv.Itoa(data)
}

func create_tokenID(info Info_tokenID) string {
	var tokenID string

	s := convert_int_to_string(info.ID)
	tokenID = create_code(s, info.Name + s, 2)

	return tokenID
}

func create_code_token(info Info_code_token, tokenID string) string {
	var code_token string

	s := convert_int_to_string(info.Role)
	code_token = create_code(s + tokenID, info.Premission, 2)

	return code_token
}


func check_token(tokenID string, code_token string) bool {
    db, err := Connect_owner()
    check_err(err)
    defer db.DB.Close()

    var code string
    err = db.DB.QueryRow("SELECT code FROM token WHERE tokenID = ?", tokenID).Scan(&code)
    if err != nil {
        if err == sql.ErrNoRows {
            return false
        }
        check_err(err)
    }

    return code == code_token
}
