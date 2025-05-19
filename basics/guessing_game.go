package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// making a Guessing Game

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1
	fmt.Println(target)

	// Welcome message
	fmt.Println("Welcome to the Guessing Game")

	fmt.Print("I have chosen a number between 1 and 100.\n Can you guess what it is? ")

	var guess int
	count := 5

	for count >= count {
		fmt.Scan(&guess)
		if target == guess {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")

		}

		count--
		fmt.Printf("You have %d attempts left.", count)
		if count < 1 {
			println("No attempt left.")
			fmt.Println("Game Over.")
			break
		}
	}

}
