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

type cow struct{}

func (c cow) Eat()   { fmt.Println("grass") }
func (c cow) Move()  { fmt.Println("walk") }
func (c cow) Speak() { fmt.Println("moo") }

type bird struct{}

func (b bird) Eat()   { fmt.Println("worms") }
func (b bird) Move()  { fmt.Println("fly") }
func (b bird) Speak() { fmt.Println("peep") }

type snake struct{}

func (s snake) Eat()   { fmt.Println("mice") }
func (s snake) Move()  { fmt.Println("slither") }
func (s snake) Speak() { fmt.Println("hiss") }

// Main function
func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("This program allows you to create a newanimal or query an existing one")
	fmt.Println("To create a new animal type 'newanimal <cow|bird|snake> <name-of-animal>'. For eg 'newanimal cow cow1'")
	fmt.Println("To query type 'query <name-of-animal> <eat|move|speak>'. For eg 'query cow1 eat'")
	fmt.Println("Press CTRL+C to exit")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "X" {
			break
		}
		fields := strings.Fields(input)
		switch fields[0] {
		case "newanimal":
			animalType := fields[1]
			name := fields[2]
			switch animalType {
			case "cow":
				animals[name] = cow{}
			case "bird":
				animals[name] = bird{}
			case "snake":
				animals[name] = snake{}
			}
			fmt.Println("Created it!")
		case "query":
			name := fields[1]
			action := fields[2]
			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found. Please create the animal first.")
				continue
			}
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		default:
			fmt.Println("Invalid command. Please try again")
		}

	}
}
