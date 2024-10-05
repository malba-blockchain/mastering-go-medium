package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		//Write a program which prompts the user to enter a floating point number
		var userPrompt string

		// Create a new reader for capturing the entire input including spaces
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("\n> Enter your request with the format: <animal> <information>")

		// Read the full string including spaces
		userPrompt, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Invalid input, please enter a valid string")
			os.Exit(0)
		}

		// Trim space before splitting to remove newline characters
		userPrompt = strings.TrimSpace(userPrompt)
		parts := strings.Split(userPrompt, " ")

		if len(parts) != 2 {
			fmt.Println("Invalid format, please enter exactly two words")
			continue
		}

		animal := strings.ToLower(parts[0])
		information := strings.ToLower(parts[1])

		var response string

		switch information {
		case "eat":
			if animal == "cow" {
				response = cow.Eat()
			}
			if animal == "bird" {
				response = bird.Eat()
			}
			if animal == "snake" {
				response = snake.Eat()
			}
		case "move":
			if animal == "cow" {
				response = cow.Move()
			}
			if animal == "bird" {
				response = bird.Move()
			}
			if animal == "snake" {
				response = snake.Move()
			}
		case "speak":
			if animal == "cow" {
				response = cow.Speak()
			}
			if animal == "bird" {
				response = bird.Speak()
			}
			if animal == "snake" {
				response = snake.Speak()
			}
		default:
			response = "Invalid information request. Please use 'eat', 'move', or 'speak'."
		}

		fmt.Println(response)
	}

}
