package main

import "fmt"

func main() {
	var word string
	fmt.Println("Enter a word: ")
	for {
		fmt.Scanln(&word)

		if word == "" {
			fmt.Println("\n Word cannot be Empty. Please enter a word.")
		} else {
			break
		}
	}

	// count the frequency of each word
	wordFrequency := make(map[string]int)

	for _, word := range word {
		wordFrequency[word]++
	}

	fmt.Println("Word Frequency: ")

	for word, frequency := range wordFrequency {
		fmt.Printf("%s: %d\n", word, frequency)
	}

}
