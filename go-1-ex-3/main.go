package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// Write to eyes.txt
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")
	// Write to dice.log
	fmt.Fprintln(os.Stderr, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go > eyes.txt 2> dice.log
}
