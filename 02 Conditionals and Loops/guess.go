package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	success := false
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I have chosen a random number between 1 and 100")
	for guesses := 10; guesses >= 0; guesses-- {
		fmt.Println("You have", guesses, "guesses left")

		fmt.Print("Make a guess :")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Your guess is too low")
		} else if guess > target {
			fmt.Println("Your guess is too high")
		} else {
			success = true
			fmt.Println("Good Handjob! You guessed it right!")
			break
		}
		if !success {
			fmt.Println("Sorry, You did'nt guess the correct number. The target was:", target)
		}
	}
}
