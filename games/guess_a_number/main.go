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

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	number := random.Intn(100)

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
