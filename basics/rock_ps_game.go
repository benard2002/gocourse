package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("WELCOME TO ROCK PAPER SCISSORS.\n==================================\n")

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	var computerMove string
	var playerMove int
	var computerScore int = 0
	var playerScore int = 0
	for {
		randHand := random.Intn(10)

		switch {
		case randHand > 6:
			computerMove = "Scissors"

		case randHand > 3:
			computerMove = "Paper"

		default:
			computerMove = "Rock"

		}
		// fmt.Println(computerMove)

		fmt.Println("\n\nComputer has chosen a hand from ROCK PAPER SCISSORS.\nSelect (1, 2, 3) for ROCK PAPER SCISSORS resp. press 0 to end game: \n")

		fmt.Scan(&playerMove)
		// fmt.Println(playerMove)

		if computerMove == "Rock" && playerMove == 2 {
			fmt.Printf("computer Chose %s. You chose Paper. You win!!", computerMove)
			playerScore++
		} else if computerMove == "Rock" && playerMove == 3 {
			fmt.Printf("computer Chose %s. You chose Scissors. You lose!!", computerMove)
			computerScore++
		} else if computerMove == "Paper" && playerMove == 1 {
			fmt.Printf("computer Chose %s. You chose Rock.You lose!!", computerMove)
			computerScore++
		} else if computerMove == "Paper" && playerMove == 3 {
			fmt.Printf("computer Chose %s. You chose Scissors. You win!!", computerMove)
			playerScore++
		} else if computerMove == "Scissors" && playerMove == 1 {
			fmt.Printf("computer Chose %s. You chose Rock. You win!!", computerMove)
			playerScore++
		} else if computerMove == "Scissors" && playerMove == 2 {
			fmt.Printf("computer Chose %s. You chose Paper. You lose!!", computerMove)
			computerScore++
		} else if playerMove <= 0 {
			fmt.Printf("CPU: %d points,  Player: %d points.", computerScore, playerScore)
			break
		} else {
			fmt.Println("Tie!")
		}

	}
}
