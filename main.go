package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}

func FirstRune(s string) rune {
	res := []rune(s)
	return res[0]
}

func splitWhiteSpaces(s string) []string {
	var str []string
	var word string
	l := len(s) - 1
	for i, v := range s {
		if i == l {
			word = word + string(v)
			str = append(str, word)
		} else if v == 32 || v == 15 || v == 10 {
			if s[i+1] == ' ' || s[i+1] == 10 {
			} else {
				str = append(str, word)
				word = ""
			}
		} else {
			word = word + string(v)
		}
	}
	return str
}

func Capitalise(s string) string {
	runes := []rune(s)

	strlen := 0
	for i := range runes {
		strlen = i + 1
	}

	for i := 0; i < strlen; i++ {
		if i != 0 && (!((runes[i-1] >= 'a' && runes[i-1] <= 'z') || (runes[i-1] >= 'A' && runes[i-1] <= 'Z'))) {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else if i == 0 {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = rune(runes[i] + 32)
			}
		}
	}
	return string(runes)
}

func RemoveTags(s []string) string {
	str := ""

	for i, tag := range s {
		if tag == "(cap," || tag == "(low," || tag == "(up," {
			s[i] = ""
			s[i+1] = ""
		} else if tag != "(up)" && tag != "(hex)" && tag != "(bin)" && tag != "(cap)" && tag != "(low)" && tag != "" {
			if i == 0 {
				str = str + tag
			} else {
				str = str + " " + tag
			}
		}
	}
	return str
}

func RemoveSpaces(s string) string {
	len := len(s) - 1
	if s[len-1] == ' ' {
		return RemoveSpaces(s[:len])
	}
	return s[:len]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func quotes(s string) string {
	str := ""
	var removeSpace bool
	for i, char := range s {
		if char == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
				removeSpace = false
			} else {
				str = str + string(char)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
			} else {
				str = str + string(char)
			}
		} else {
			str = str + string(char)
		}
	}
	return str
}

func main() {
	sample := os.Args[1]
	data, err := os.ReadFile(sample)
	check(err)
	input := string(data)
	result := splitWhiteSpaces(input)

	for i, v := range result {
		// // converts number before "(hex)" before to decimal
		if Compare(v, "(hex)") == 0 {
			j, _ := strconv.ParseInt(result[i-1], 16, 64)
			result[i-1] = fmt.Sprint(j)
		}
		// converts number before "(bin)" before to decimal
		if Compare(v, "(bin)") == 0 {
			j, _ := strconv.ParseInt(result[i-1], 2, 64)
			result[i-1] = fmt.Sprint(j)
		}
		// convert word before to lowercase
		if Compare(v, "(low)") == 0 {
			result[i-1] = strings.ToLower(result[i-1])
		}

		//converts the number of words before to lowercase
		if Compare(v, "(low,") == 0 {
			result[i-1] = strings.ToLower(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)

			for j := 1; j <= nu; j++ {
				result[i-j] = strings.ToLower(result[i-j])
			}
		}

		//convert word before to uppercase
		if Compare(v, "(up)") == 0 {
			result[i-1] = strings.ToUpper(result[i-1])
		}

		// converts the number of words before to uppercase
		if Compare(v, "(up,") == 0 {
			result[i-1] = strings.ToLower(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)

			for j := 1; j <= nu; j++ {
				result[i-j] = strings.ToUpper(result[i-j])
			}
		}

		//converts the word before to capitalise
		if Compare(v, "(cap)") == 0 {
			result[i-1] = Capitalise(result[i-1])

		}
		// converts the number of words before to capitalise
		if Compare(v, "(cap,") == 0 {
			result[i-1] = Capitalise(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)

			for j := 1; j <= nu; j++ {
				result[i-j] = Capitalise(result[i-j])
			}
		}

		vowels := "aeiouh"
		if i > 0 {
			firstRune := string(FirstRune(result[i]))
			// converts 'a' into 'an' when the next word begins with a vowel or 'h'.
			if Compare(result[i-1], "a") == 0 && strings.Contains(vowels, firstRune) {
				result[i-1] = "an"
			}
			if Compare(result[i-1], "A") == 0 && strings.Contains(vowels, firstRune) {
				result[i-1] = "An"
			}
		}
	}

	noTagResult := RemoveTags(result)
	result2 := splitWhiteSpaces(noTagResult)
	str := ""
	for _, word := range result2 {
		str = str + word + " "
	}

	// Removing the string
	str = RemoveSpaces(str)
	word := ""
	for i, char := range str {
		if i == len(str)-1 {
			if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
				if str[i-1] == ' ' {
					word = word[:len(word)-1]
					word = word + string(char)
				} else {
					word = word + string(char)
				}
			} else {
				word = word + string(char)
			}
		} else if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
			if str[i-1] == ' ' {
				word = word[:len(word)-1]
				word = word + string(char)
			} else {
				word = word + string(char)
			}
			if str[i+1] != ' ' && str[i+1] != '.' && str[i+1] != ',' && str[i+1] != '!' && str[i+1] != '?' && str[i+1] != ';' && str[i+1] != ':' {
				word = word + " "
			}
		} else {
			word = word + string(char)
		}
	}

	//Removing quotes
	word = quotes(word)

	//fmt.Println(word)
	err = os.WriteFile("result.txt", []byte(word), 0755)
	if err != nil {
		fmt.Println(err)
	}
}
