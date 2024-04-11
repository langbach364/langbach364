package main

import "fmt"

func check(s1 string, s2 string) bool {
	return s1 == s2
}

type s struct {
	s1 string
	s2 string
}

func main() {
	var s1 s
	s1.s1 = "Hello"
	s1.s2 = " World"
	s2 := s1

	fmt.Println("check encode:", check(encode_data(s1.s1, s1.s2, 2), encode_data(s2.s1, s2.s2, 2)))
	fmt.Println("algorithm after encode:",encode_data(s1.s1, s1.s2, 2))
}
