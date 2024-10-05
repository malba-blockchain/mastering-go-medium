package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct {
	eatenFood        string
	locomationMethod string
	spokenSound      string
}

func (this *Cow) Eat() string {
	return this.eatenFood
}

func (this *Cow) Move() string {
	return this.locomationMethod
}
func (this *Cow) Speak() string {
	return this.spokenSound
}

type Bird struct {
	eatenFood        string
	locomationMethod string
	spokenSound      string
}

func (this *Bird) Eat() string {
	return this.eatenFood
}

func (this *Bird) Move() string {
	return this.locomationMethod
}
func (this *Bird) Speak() string {
	return this.spokenSound
}

type Snake struct {
	eatenFood        string
	locomationMethod string
	spokenSound      string
}

func (this *Snake) Eat() string {
	return this.eatenFood
}

func (this *Snake) Move() string {
	return this.locomationMethod
}
func (this *Snake) Speak() string {
	return this.spokenSound
}

func NewAnimal(typ string) Animal {
	switch strings.ToLower(typ) {
	case "cow":
		return &Cow{
			eatenFood:        "grass",
			locomationMethod: "walk",
			spokenSound:      "moo",
		}
	case "bird":
		return &Bird{
			eatenFood:        "worms",
			locomationMethod: "fly",
			spokenSound:      "peep",
		}
	case "snake":
		return &Snake{
			eatenFood:        "mice",
			locomationMethod: "slither",
			spokenSound:      "hsss ",
		}
	default:
		panic("Unsupported typ value")
	}

}

func determineMethod(animl Animal, methodName string) string {
	switch strings.ToLower(methodName) {
	case "eat":
		return animl.Eat()
	case "move":
		return animl.Move()
	case "speak":
		return animl.Speak()
	default:
		panic("Unsupported method value")
	}
}
func main() {
	mapAnimals := make(map[string]Animal)

	consoleReader := bufio.NewScanner(os.Stdin)

	for consoleReader.Scan() {
		inputs := strings.Split(consoleReader.Text(), " ")

		command := inputs[0]
		switch command {
		//newanimal <arbitrary-name> <typ-name>
		case "newanimal":
			animlName := inputs[1]
			animlType := inputs[2]

			mapAnimals[animlName] = NewAnimal(animlType)
		//query <name of animal> <method-name>
		case "query":
			animlName := inputs[1]
			animlAction := inputs[2]
			animl, ok := mapAnimals[animlName]

			if !ok {
				fmt.Printf("Animal %s not found", animlName)
				continue
			}
			actionResult := determineMethod(animl, animlAction)

			fmt.Printf("Animal: %s Method: %s Result %s \n",
				animlName, animlAction, actionResult)
		default:
			panic("Not supported command ")
		}
	}
}
