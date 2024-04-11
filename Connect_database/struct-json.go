package main

import ()

type query struct {
	Query string `json:"query"`
}

type send_email struct {
	Email_sender    string `json:"email_sender"`
	Password_sender string `json:"password"`
	Email_recevier  string `json:"email_receiver"`
}

type object struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Role int `json:"role"`
	Premission string `json:"premission"`
}

type check_code struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type Info_tokenID struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type Info_code_token struct {
	Role int `json:"role"`
	Premission string `json:"premission"`
}
