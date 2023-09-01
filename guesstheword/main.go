package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println(getHangmanDrawing(6))
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

func printGameState(targetWord string, guessedLetters map[rune]bool) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
}

func getWordGuessingProgress(
	targetWord string,
	guessedLetters map[rune]bool,
) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}

		result += " "
	}

	return result
}

func initializeGuessedWord(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(
		fmt.Sprintf("states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}

	return string(data)
}
