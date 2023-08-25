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
	targetWord := dictionary[rand.Intn(len(dictionary))]
	fmt.Println(targetWord)
}
