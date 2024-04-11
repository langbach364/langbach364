package main

import (
)

func check_number_code(data *int) int {
	if check_rune_1(*data) && check_rune_2(*data) {
		return *data
	}
	for *data < 33 || (*data >= 48 && *data <= 57) {
		*data = *data + 32
	}
	*data = check_size_number(*data)
	return *data
}


func insert_string_code(mixing []int, HashData []string) []rune {
	result := make([]rune, 0)
	insert_rules := insert_rules(mixing, len(HashData))

	for i := 0; i < len(mixing); i++ {
		pointer := insert_rules[i]
		number := convert_string_to_int(HashData[pointer])
		check_number_code(&number)
		char := convert_int_to_ascii(number)
		result = append(result, char)
	}
	return result
}

func get_value_excessive_code(SizeMerge int, current int) int {
	return SizeMerge - current
}

func check_excessive_data_code(SizeMerge int, current int, hash int) bool {
	return current+hash <= SizeMerge
}

func hash_data_code(merge string, SizeHash int, SaveExcessive *int) []string {
	var result []string
	for i := 0; i < len(merge); i += SizeHash {
		if !check_excessive_data_code(len(merge), i, SizeHash) {
			*SaveExcessive = get_value_excessive_code(len(merge), i)
			break
		}
		str := i + SizeHash
		result = append(result, merge[i:str])
	}
	return result
}

func data_mixing_code(mixing []int, HashData []string) string {
	var result string

	pointer := 0
	for i := 0; i < len(mixing); i++ {
		x := mixing[pointer]
		swap(&HashData[i], &HashData[x])
		pointer++
	}
	chars := insert_string_code(mixing, HashData)
	for i := 0; i < len(HashData); i++ {
		result += HashData[i] + string(chars[i])
	}
	return result
}

func create_code(dataGmail string, dataPassword string, size_hash int) string {
	if size_hash < 2 {
		size_hash = 2
	}
	asscii_gmail := convert_string_to_ascii(dataGmail)
	asscii_password := convert_string_to_ascii(dataPassword)
	merge := merge_string(asscii_gmail, asscii_password)
	save_excessive := 0
	hashData := hash_data_code(merge, size_hash, &save_excessive)
	mixing := mixing_rules(merge, len(hashData))
	dataMixing := data_mixing_code(mixing, hashData)

	return dataMixing
}

///////////////////////////////////////////////////////////////////////////////////////////
