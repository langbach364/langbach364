package main

import (
	"fmt"
)

///////////////////////////////////////////////////////////
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
	if len(data) > current_position {
		return true
	}
	return false
}

///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////
// Ghép ascii của gmail và password 
func check_size(data []rune, pointer int) bool {
	if pointer < len(data) {
		return true
	}
	return false
}

func check_merge(dataGmail []rune, pointer int) bool {
	if check_size(dataGmail, pointer) {
		return true
	}
	return false
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

///////////////////////////////////////////////////////////////////////////////////////////
// Tạo quy tắc trộn
func check_symbol(x int, size int) bool {
	if size > x {
		return true
	}
	return false
}

func check_length(x int, size int) bool {
	if x > size {
		return true
	}
	return false
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
		*x = size - 1
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
			convert_symbol(&sum, SizeHash)
		}

		result[i] = sum
		count--
	}
	return result
}

///////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////
// Băm dữ liệu ra và lấy phần thừa của việc băm rồi ghép lại
func get_value_excessive(SizeMerge int, current int) int {
	return SizeMerge - current
}

func check_excessive_data(SizeMerge int, current int, hash int) bool {
	if current+hash > SizeMerge {
		return false
	}
	return true
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
		swap(&HashData[i], &HashData[mixing[pointer]])
		pointer++
	}
	for i := 0; i < len(HashData); i++ {
		result += HashData[i]
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

///////////////////////////////////////////////////////////////////////////////////////////
// Bắt đầu việc mã hóa(tạo nhiều đối tượng cho dễ debug hơn)
func encode_data(dataGmail string, dataPassword string, size_hash int) string {
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

func check_function(check1 string, check2 string) bool {
	if check1 != check2 {
		return false
	}
	return true
}

func main() {
	check1 := encode_data("09123123123", "1234laksdjbv5", 2)
	check2 := encode_data("09123123123", "1234laksdjbv5", 2)

	fmt.Println("TRUE LÀ CODE CHẠY ĐÚNG, FALSE LÀ CODE CHẠY SAI:", check_function(check1, check2))
	fmt.Println("check1: ", check1)
	fmt.Println("check2: ", check2)
}
