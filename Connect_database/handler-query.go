package main

import ()

// Hàm này kiểm tra cấu trúc của câu truy vấn
// This is function to check the structure of the query
func Structure_query(query string, Struct string) bool {
	words := split_words(query)
	if !check_query(words) {
		return false
	}
	if !check_condition_allowed(words) {
		return false
	}
	if !check_sructure_query(words, Struct) {
		return false
	}
	return true
}