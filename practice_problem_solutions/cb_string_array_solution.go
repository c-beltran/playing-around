// Problem 1 solution
package main

import (
	"bytes"
	"fmt"
	"strings"
)

var Vowels = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
}

func main() {
	// read a sentence and print the longest word
	// and the word that contains the most vowels

	words := strings.Split("the three musketeers saved the day from a fierce dragon", " ")

	length := len(words)
	var longestWord string

	for i := 0; i < length; i++ {
		if len(words[i]) > len(longestWord) {
			longestWord = words[i]
		}
	}

	fmt.Println("longest word: ", longestWord)

	//find the word with the longest vowel and vowel count

	// words = []string{"bat", "bola"}

	winningWord := [2]string{}

	for _, w := range words {
		v := getVowels(w)

		if len(winningWord[1]) < len(v) {
			winningWord[0] = w
			winningWord[1] = v

		}
	}

	fmt.Println(winningWord)

	// cannot have repeated vowels

}

// create a helper function to return the vowels
func getVowels(word string) string {
	if len(word) == 0 {
		return ""
	}

	var b bytes.Buffer

	for _, l := range word {
		if Vowels[string(l)] {

			if !strings.ContainsAny(string(l), b.String()) {

				b.WriteString(string(l))

			}
		}
	}

	return b.String()
}
