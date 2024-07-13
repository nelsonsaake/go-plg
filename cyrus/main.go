package main

import (
	"fmt"
	"strconv"
	"strings"
)

// chunk splits a string into chunks of size chunkSize
func chunk(s string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

func decode(input string) string {

	var (
		output = ""
		chunks = chunk(input, 2)
	)

	for _, chunk := range chunks {
		num, err := strconv.Atoi(chunk)
		if err != nil {
			fmt.Println("Error converting string to number:", err)
			return ""
		}
		output += string(rune(num + 'a' - 1))
	}

	return output
}

func encode(input string) string {

	var (
		output = ""
	)

	for _, char := range input {
		num := int(char - 'a' + 1)
		strNum := strconv.Itoa(num)
		if num < 10 {
			strNum = "0" + strNum // Ensure two digits
		}
		output += strNum
	}

	return output
}

func main() {

	v := "0114092001 16091411"
	vs := strings.Split(v, " ")

	for _, v := range vs {
		fmt.Print(decode(v))
	}
	fmt.Println()

	fmt.Println(encode("abc"))
}
