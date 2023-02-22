/*
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake.

Each animal can eat, move, and speak.

The user can issue a request to find out one of three things about an animal:

1) the food that it eats,
2) its method of locomotion, and
3) the sound it makes when it speaks.

The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal    | Food eaten      |	Locomotion method     |	Spoken sound
=========================================================
cow       |	grass	        | walk	                  |  moo
bird	  | worms	        | fly	                  |  peep
snake	  | mice	        | slither	              |  hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request,
and prints out a new prompt. Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:food, locomotion, and noise,
all of which are strings. Make three methods called Eat(), Move(), and Speak().
The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.

*/

package main

import (
	"errors"
	"fmt"
)

// Struct about Animals
type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

// Declare the animals
var cow = Animal{"cow", "grass", "walk", "moo"}
var bird = Animal{"bird", "worms", "fly", "peep"}
var snake = Animal{"snake", "mice", "slither", "hsss"}

// Eat method print the food of the animal
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move method print the locomotion of the animal
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak method print the noise of the animal
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

// assign the animal to the animal name or return error

func assignAnimal(animalName string) (animal Animal, err error) {
	switch animalName {
	case "cow":
		animal = cow
	case "bird":
		animal = bird
	case "snake":
		animal = snake
	default:
		return animal, errors.New("invalid animal name")
	}
	return animal, nil
}

// assign the action to the action name
func assignAction(actionName string, animal Animal) {
	switch actionName {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Please enter a valid action name")
	}
}

func main() {
	var animalName, actionName string
	fmt.Println("Please type '[cow|bird|snake] [eat|move|speak]' after" +
		"the prompt \">\" to discover animals. Press CTRL+C to exit.")
	for {
		fmt.Print("> ")
		fmt.Scan(&animalName)
		animal, err := assignAnimal(animalName)
		if err != nil {
			fmt.Println(err)
			continue // skip the rest of the loop
		}
		fmt.Scan(&actionName)
		assignAction(actionName, animal)
	}
}
