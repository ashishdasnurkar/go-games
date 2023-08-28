package main

import (
	"fmt"
	"math/rand"
	"time"
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
	guessedLetters := map[rune]bool{}
	guessedLetters[rune(targetWord[0])] = true
	guessedLetters[rune(targetWord[len(targetWord)-1])] = true
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
			fmt.Print(" _ ")
		}
	}
	fmt.Println()
}
