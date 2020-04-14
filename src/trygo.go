package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}
func (a *Animal) Move() string {
	return a.locomotion
}
func (a *Animal) Speak() string {
	return a.noise
}
func main() {
	cowVal := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	birdVal := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	snakeVal := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
	fmt.Println("Enter request of type `animalName infoRequested`")
	var animalType string
	var animalInfo string
	var returnType Animal
	var returnVal string
	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s", &animalType, &animalInfo)
		switch animalType {
		case "cow":
			returnType = cowVal
		case "bird":
			returnType = birdVal
		case "snake":
			returnType = snakeVal
		default:
			fmt.Println("Invalid animal type entered")
			continue
		}
		switch animalInfo {
		case "eat":
			returnVal = returnType.Eat()
		case "move":
			returnVal = returnType.Move()
		case "speak":
			returnVal = returnType.Speak()
		default:
			fmt.Println("Invalid property entered")
			continue
		}
		fmt.Println(returnVal)
	}

}
