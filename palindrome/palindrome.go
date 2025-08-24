package palindrome

import "fmt"

func getString() (str string) {
	fmt.Println("write string : ")
	_, err := fmt.Scanf("%v", &str)

	if err != nil {
		return
	}
	return
}

func Palindrome() bool {
	s := getString()
	runes := []rune(s)
	start := 0
	end := len(runes) - 1

	for start < end {
		if runes[start] != runes[end] {
			return false
		}
		start++
		end--
	}
	return true
}
