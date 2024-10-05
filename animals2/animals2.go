package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() { fmt.Println("grass") }

func (c Cow) Move() { fmt.Println("walk") }

func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (c Bird) Eat() { fmt.Println("worms") }

func (c Bird) Move() { fmt.Println("fly") }

func (c Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (c Snake) Eat() { fmt.Println("mice") }

func (c Snake) Move() { fmt.Println("slither") }

func (c Snake) Speak() { fmt.Println("hsss") }

func main() {

	animals := make(map[string]Animal)

	for {
		//Write a program which prompts the user to enter a floating point number
		var userPrompt string

		// Create a new reader for capturing the entire input including spaces
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("\n> Enter your request with the format:\n [newanimal] <animalName> <animalType> or [query] <animalName> <information>")

		// Read the full string including spaces
		userPrompt, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("> Invalid input, please enter a valid string.")
			continue
		}

		// Trim space before splitting to remove newline characters
		userPrompt = strings.TrimSpace(userPrompt)
		parts := strings.Split(userPrompt, " ")

		if len(parts) != 3 {
			fmt.Println("> Invalid format, please enter your command with 3 strings.")
			continue
		}

		command := strings.ToLower(parts[0])

		if command == "newanimal" {

			animalName := strings.ToLower(parts[1])

			animalType := strings.ToLower(parts[2])

			var animal Animal

			switch animalType {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Println("> Invalid animal type. Use 'cow', 'bird', or 'snake'.")
				continue
			}

			animals[animalName] = animal

			fmt.Println("> Created it!")

		} else if command == "query" {

			animalName := strings.ToLower(parts[1])

			animalInformation := strings.ToLower(parts[2])

			animal, exists := animals[animalName]

			if !exists {
				fmt.Println("> Error. Animal name doesn't exists.")
				continue
			}

			switch animalInformation {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("> Invalid information request.")
				continue
			}

		} else {
			fmt.Println("> Invalid value, please enter a valid command.")
			continue
		}
	}
}
