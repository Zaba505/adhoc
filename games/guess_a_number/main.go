package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1) Create a random number to guess
// 2) Get users' guess
// 3) Check users' guess
// 4) Tell user: too high, too low, or correct

func randInt(r *rand.Rand, low, high int) int {
	return r.Intn(high-low) + low
}

func runGame(r *rand.Rand, low, high int) {
	number := randInt(r, low, high)

	var guess int
	for {
		fmt.Print("Guess a number: ")
		fmt.Scanf("%d", &guess)

		if guess == number {
			fmt.Println("You're correct!")
			break
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess < number {
			fmt.Println("Too low")
		}
	}
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	runGame(r, 0, 100)
}
