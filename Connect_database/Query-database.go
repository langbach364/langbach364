package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"errors"
)

func validate_And_Create_Token(query query, body []byte, command string) error {
    if !Structure_query(query.Query, command) {
        return errors.New("query is not allowed")
    }

    var object object
    err := json.Unmarshal(body, &object)
    if err != nil {
        return err
    }

    var Info_tokenID Info_tokenID
    Info_tokenID.ID = object.ID
    Info_tokenID.Name = object.Name

    var info_code_token Info_code_token
    info_code_token.Role = object.Role
    info_code_token.Premission = object.Premission

    tokenID := create_tokenID(Info_tokenID)
    code_token := create_code_token(info_code_token, tokenID)

    if !check_token(tokenID, code_token) {
        return errors.New("token is not correct")
    }

    return nil
}

func insert_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				err = validate_And_Create_Token(query, body, "insert")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusMethodNotAllowed)
                    return
                }

				_, err = dbInfo.DB.Exec(query.Query)
				check_err(err)
				json.NewEncoder(w).Encode("true")
			}
		}
	}
}

func select_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				err = validate_And_Create_Token(query, body, "select")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusMethodNotAllowed)
                    return
                }

				rows, err := dbInfo.DB.Query(query.Query)
				check_err(err)
				defer rows.Close()

				var data []map[string]interface{}

				columns, err := rows.Columns()
				check_err(err)

				for rows.Next() {
					values := make([]interface{}, len(columns))
					valuePtrs := make([]interface{}, len(columns))
					for i := range values {
						valuePtrs[i] = &values[i]
					}

					err := rows.Scan(valuePtrs...)
					check_err(err)

					row := make(map[string]interface{})
					for i, column := range columns {
						row[column] = values[i]
					}

					data = append(data, row)
				}
				json.NewEncoder(w).Encode(data)
			}
		default:
			http.Error(w, "Method is not used", http.StatusMethodNotAllowed)
		}
	}
}

func delete_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				err = validate_And_Create_Token(query, body, "delete")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusMethodNotAllowed)
                    return
                }

				_, err = dbInfo.DB.Exec(query.Query)
				check_err(err)

				json.NewEncoder(w).Encode("true")
			}
		default:
			http.Error(w, "Method is not used", http.StatusMethodNotAllowed)
		}
	}
}

func update_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				err = validate_And_Create_Token(query, body, "update")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusMethodNotAllowed)
                    return
                }

				_, err = dbInfo.DB.Exec(query.Query)
				check_err(err)

				json.NewEncoder(w).Encode("true")
			}
		}
	}
}
func send_code() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var send_token send_email
				err = json.Unmarshal(body, &send_token)
				check_err(err)

				Send_Email(send_token.Email_sender, send_token.Password_sender, send_token.Email_recevier)
				json.NewEncoder(w).Encode("true")
			}
		default:
			http.Error(w, "Method is not used", http.StatusMethodNotAllowed)
		}
	}
}

func session_account() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var code check_code
				err = json.Unmarshal(body, &code)
				check_err(err)

				if verify_code(code.Email, code.Code) {
					json.NewEncoder(w).Encode("true")
				} else {
					json.NewEncoder(w).Encode("false")
				}
			}
		default:
			http.Error(w, "Method is not used", http.StatusMethodNotAllowed)
		}
	}
}
