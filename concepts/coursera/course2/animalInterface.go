/*

Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create
a new animal of one of the three types, or the user can request information about an animal that he/she has
already created. Each animal has a unique name, defined by the user. Note that the user can define animals of
 a chosen type, but the types of animals are restricted to either cow, bird, or snake.
 The following table contains the three types of animals and their associated data.
Animal    | Food eaten    |	Locomotion method |	Spoken sound
=========================================================
cow       |	grass	      | walk	          |  moo
bird	  | worms	      | fly	              |  peep
snake	  | mice	      | slither	          |  hsss


Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
 program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
  Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new
animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal.
The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain
the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food,
the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three
types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird,
and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type.
Your program should call the appropriate method when the user issues a query command.
*/

package main

import "fmt"

// Interface for animals
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Struct for animals
type AnimalStruct struct {
	food       string
	locomotion string
	noise      string
}

// Method for eating with name
func (a *AnimalStruct) Eat() {
	fmt.Println(a.food)
}

// Method for moving
func (a *AnimalStruct) Move() {
	fmt.Println(a.locomotion)
}

// Method for speaking
func (a *AnimalStruct) Speak() {
	fmt.Println(a.noise)
}

// Make coe, bird and snake
var cow = AnimalStruct{"grass", "walk", "moo"}
var bird = AnimalStruct{"worms", "fly", "peep"}
var snake = AnimalStruct{"mice", "slither", "hsss"}

// Map for animals
var animalsMap = map[string]AnimalStruct{
	"cow":   cow,
	"bird":  bird,
	"snake": snake,
}

// Main function
func main() {
	var command, animalReq1, animalReq2 string
	fmt.Println("This program allows you to create a newanimal or query an exsisting one")
	fmt.Println("To create a new animal type 'newanimal <cow|bird|snake> <name-of-animal>'. For eg 'newanimal cow cow1'")
	fmt.Println("To query type 'query <name-of-animal> <eat|move|speak>'. For eg 'quet cow1 eat'")

	for {
		fmt.Printf(">")
		fmt.Scan(&command, &animalReq1, &animalReq2)
		if command == "newanimal" {
			animalsMap[animalReq2] = animalsMap[animalReq1]
			fmt.Println("Created it!")
		}
		if command == "query" {
			genralAni := animalsMap[animalReq1]
			switch animalReq2 {
			case "eat":
				genralAni.Eat()
			case "move":
				genralAni.Move()
			case "speak":
				genralAni.Speak()
			}
		}
	}
}
