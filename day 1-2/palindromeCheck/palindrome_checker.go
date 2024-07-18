package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkPalindrome(sentence string) bool {
	left, right := 0, len(sentence)-1

	for left < right {
		if sentence[left] != sentence[right] {
			return false

		}
		left++
		right--

	}
	return true

}

func main() {

	// intialize the bufio
	reader := bufio.NewReader(os.Stdin)

	var sentence string
	// accept the string as input

	fmt.Println("Input a String: ")
	for {
		line, err := reader.ReadString('\n')
		sentence = line

		if err != nil {
			fmt.Println("Please Input a valid String. ")
		} else {
			break
		}
	}

	if checkPalindrome(strings.TrimSpace(sentence)) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a Palindrome")
	}

}
