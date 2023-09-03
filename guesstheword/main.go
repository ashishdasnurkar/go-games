package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)
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
	hangmanState := 0
	for !isWordGuessed(targetWord, guessedLetters) && !isHangmanComplete(hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()

		if len(input) != 1 {
			fmt.Println("Please input a single letter only...")
			continue
		}
		letter := rune(input[0])
		if isCorrectGuess(targetWord, letter) {
			guessedLetters[rune(input[0])] = true
		} else {
			hangmanState++
		}
	}
	printGameState(targetWord, guessedLetters, hangmanState)
	fmt.Print("Game over ... ")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You won!!!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You lose :(")
	}
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
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

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 7
}
func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(
		fmt.Sprintf("states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}

	return string(data)
}

func readInput() string {
	fmt.Print("> ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}

func isCorrectGuess(targeWord string, letter rune) bool {
	return strings.ContainsRune(targeWord, letter)
}
