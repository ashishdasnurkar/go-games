package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

var dictionary = []string{
	"Swashank",
	"Phenomenon",
	"Sideways",
	"Departed",
	"parker",
	"Memento",
	"Gladiator",
	"Taken",
	"Jumper",
	"Fury",
	"Inglorious",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := getRandomWord()
	guessedLetters := initializeGuessedWord(targetWord)
	printGameState(targetWord, guessedLetters)
	fmt.Println(targetWord)
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

func printGameState(targetWord string, guessedLeters map[rune]bool) {
	for _, ch := range targetWord {
		if guessedLeters[ch] == true {
			fmt.Printf("%c", ch)
		} else {
			fmt.Print("_")
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

func initializeGuessedWord(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}
