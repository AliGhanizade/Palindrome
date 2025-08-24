package main

import (
	"Palindrome/palindrome"
	"fmt"
)

func main() {
	check := palindrome.Palindrome()
	if !check {
		fmt.Print("False")
		return
	}
	fmt.Print("True")

}
