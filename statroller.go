package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
type stats struct {
	stats [6]int
}
*/

func getChoice() {
	var choice, dice int
	fmt.Println("\n1. Standard\n2. Classic\n3. Heroic\nChoose a roll type: ")

	fmt.Scan(&choice)

	if choice == 1 {
		dice = 4
		statResults(dice)
	}
	if choice == 2 {
		dice = 3
		statResults(dice)

	}
	if choice == 3 {
		dice = 2
		statResults(dice)
	}
}

func statResults(dice int) {
	var stats [6]int

	//populate array with 6 stats
	for x := 0; x <= 5; x++ {
		stats[x] = statCalc(dice)
	}

	fmt.Printf("\n%d, %d, %d, %d, %d, %d\n", stats[0], stats[1], stats[2], stats[3], stats[4], stats[5])
}

func statCalc(dice int) int {
	var total, small int

	//total = small = diceRoll()
	small = diceRoll()
	total = small

	for i := 1; i < dice; i++ {
		roll := diceRoll()
		total += roll
		if roll < small {
			small = roll
		}
	}
	//Heroic
	if dice == 2 {
		total += 6
	}
	//Standard Roll
	if dice == 4 {
		fmt.Printf("\nSmallest dice roll should be %d", small)
		fmt.Printf("\nTotal before removal of smallest dice %d", total)
		total -= small
	}
	return total
}

func diceRoll() int {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	return rand.Intn(6) + 1
}

func main() {
	fmt.Println("Welcome to the Stat Roller!")
	fmt.Println("This Stat Roller can use the #@$&#$*#&")

	getChoice()
}
