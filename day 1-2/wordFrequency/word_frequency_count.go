package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// bufio is used when we want to read in a full line
	reader := bufio.NewReader(os.Stdin)
	var sentence string
	fmt.Println("Enter a sentence: ")
	for {
		line, err := reader.ReadString('\n')
		sentence = line

		if sentence == "" || err != nil {
			fmt.Println("\n sentence cannot be Empty. Please enter a sentence.")
		} else {
			break
		}
	}

	// count the frequency of each sentence
	wordFrequency := make(map[string]int)

	for _, word := range strings.Split(sentence, " ") {
		wordFrequency[strings.TrimSpace(word)]++
	}

	fmt.Println("Word Frequency: ")

	for word, frequency := range wordFrequency {
		fmt.Printf("%s: %d\n", word, frequency)
	}

}
