package helper

import "strconv"

func AddSeparator(bilangan int) string {
	var numberString = strconv.Itoa(bilangan)

	var result string

	for i, char := range numberString {
		if i > 0 && (len(numberString)-i)%3 == 0 {
			result += "."
		}
		result += string(char)
	}

	return result
}

func AddSeparatorInt64(bilangan int64) string {
	var numberString = strconv.FormatInt(bilangan, 10)

	var result string

	for i, char := range numberString {
		if i > 0 && (len(numberString)-i)%3 == 0 {
			result += "."
		}
		result += string(char)
	}

	return result
}
