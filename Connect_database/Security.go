package main

import (
	"fmt"
)

// /////////////////////////////////////////////////////////
// Hàm phụ hỗ trợ xử lý
func convert_string_to_ascii(data string) []rune {
	runes := []rune(data)
	ascii := make([]int64, len(runes))
	result := make([]rune, len(ascii))

	for i, r := range runes {
		ascii[i] = int64(r)
		result[i] = rune(ascii[i])
	}
	return result
}

func swap(a *string, b *string) {
	temp := *a
	*a = *b
	*b = temp
}

func get_excessive(current_position int, data []rune) []int {
	result := make([]int, len(data)-current_position)
	var pointer int = 0
	for i := current_position; i < len(data); i++ {
		result[pointer] = int(data[i])
		pointer++
	}
	return result
}

func check_excessive(data []rune, current_position int) bool {
	return len(data) > current_position
}

///////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////////////////////////////
// Ghép ascii của gmail và password
func check_size(data []rune, pointer int) bool {
	return pointer < len(data)
}

func check_merge(dataGmail []rune, pointer int) bool {
	return check_size(dataGmail, pointer)
}

func merge_string(dataGmail []rune, dataPassword []rune) string {
	var merge string

	var SavePointer int
	for pointer := range dataPassword {
		if check_merge(dataGmail, pointer) {
			merge += fmt.Sprintf("%d%d", dataPassword[pointer], dataGmail[pointer])
		} else {
			merge += fmt.Sprintf("%d", dataPassword[pointer])
		}
		SavePointer = pointer
	}

	if check_excessive(dataGmail, SavePointer) {
		excess := get_excessive(SavePointer, dataGmail)
		for i := range excess {
			merge += fmt.Sprintf("%d", excess[i])
		}
	}
	return merge
}

///////////////////////////////////////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////////////////////////////////////
// Tạo quy tắc trộn
func check_symbol(x int, size int) bool {
	return size > x
}

func check_length(x int, size int) bool {
	return x > size
}

func convert_symbol(x *int, size int) int {
	if *x > size && *x < 10 {
		*x = 0
		return *x
	}
	slice1 := *x % 10
	slice2 := *x / 10
	*x = slice1 + slice2

	if check_length(*x, size) {
		*x = size
	}
	return *x
}

func convert_rune_to_int(char rune) int {
	return int(char - '0')
}

func mixing_rules(merge string, SizeHash int) []int {
	result := make([]int, SizeHash)

	count := SizeHash

	for i := 0; i < SizeHash; i++ {
		if count == 0 {
			break
		}
		sum := convert_rune_to_int(rune(merge[i])) + convert_rune_to_int(rune(merge[i+1]))

		if !check_symbol(sum, SizeHash) {
			convert_symbol(&sum, SizeHash - 1)
		}

		result[i] = sum
		count--
	}
	return result
}

///////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////
// Tạo quy tắc chèn và các ký tự

func convert_string_to_int(data string) int {
	var result int
	for _, char := range data {
		result = result*10 + int(char-'0')
	}
	return result
}

func check_rune_1(data int) bool {
	return data > 32 && data < 127
}

func check_rune_2(data int) bool {
	return data < 47 || data > 57
}

func check_size_number(data int) int {
	if data > 126 {
		data = 64
	}
	return data
}

func check_number(data *int) int {
	if check_rune_1(*data) && check_rune_2(*data) {
		return *data
	}
	for *data < 32 || (*data >= 48 && *data <= 57) {
		*data = *data + 32
	}
	*data = check_size_number(*data)
	return *data
}

func convert_int_to_ascii(data int) rune {
	return rune(data)
}

func insert_rules(mixing []int, size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = size - mixing[i]
		if result[i] >= size {
			result[i] = size - 1
		}
	}
	return result
}

func insert_strings(mixing []int, HashData []string) []rune {
	result := make([]rune, 0)
	insert_rules := insert_rules(mixing, len(HashData))

	for i := 0; i < len(mixing); i++ {
		pointer := insert_rules[i]
		number := convert_string_to_int(HashData[pointer])
		check_number(&number)
		char := convert_int_to_ascii(number)
		result = append(result, char)
	}
	return result
}

// /////////////////////////////////////////////////////////////////////////////////////////
// /////////////////////////////////////////////////////////////////////////////////////////
// Băm dữ liệu ra và lấy phần thừa của việc băm rồi ghép lại
func get_value_excessive(SizeMerge int, current int) int {
	return SizeMerge - current
}

func check_excessive_data(SizeMerge int, current int, hash int) bool {
	return current+hash <= SizeMerge
}

func hash_data(merge string, SizeHash int, SaveExcessive *int) []string {
	var result []string
	for i := 0; i < len(merge); i += SizeHash {
		if !check_excessive_data(len(merge), i, SizeHash) {
			*SaveExcessive = get_value_excessive(len(merge), i)
			break
		}
		str := i + SizeHash
		result = append(result, merge[i:str])
	}
	return result
}

func data_mixing(mixing []int, HashData []string) string {
	var result string

	pointer := 0
	for i := 0; i < len(mixing); i++ {
		x := mixing[pointer]
		swap(&HashData[i], &HashData[x])
		pointer++
	}
	chars := insert_strings(mixing, HashData)
	for i := 0; i < len(HashData); i++ {
		result += HashData[i] + string(chars[i])
	}
	return result
}

func combine_excessive(merge string, save_excessive int) string {
	var result string
	for i := save_excessive; i < len(merge); i++ {
		result += string(merge[i])
	}
	return result
}

///////////////////////////////////////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////////////////////////////////////
// Bắt đầu việc mã hóa(tạo nhiều đối tượng cho dễ debug hơn)
func encode_data(dataGmail string, dataPassword string, size_hash int) string {
	if size_hash < 2 {
		size_hash = 2
	}
	asscii_gmail := convert_string_to_ascii(dataGmail)
	asscii_password := convert_string_to_ascii(dataPassword)
	merge := merge_string(asscii_gmail, asscii_password)
	save_excessive := 0
	hashData := hash_data(merge, size_hash, &save_excessive)
	mixing := mixing_rules(merge, len(hashData))
	excessive := combine_excessive(merge, save_excessive)
	dataMixing := data_mixing(mixing, hashData) + excessive

	return dataMixing
}

///////////////////////////////////////////////////////////////////////////////////////////
