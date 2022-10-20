package main

import "fmt"

func main() {

	// Arrays

	var ages = [3]int{15, 24, 65}

	names := [4]string{"Nathan", "Addy", "Elliot", "Angela"}

	names[1] = "Adrienne"

	fmt.Println(ages, len(ages))

	fmt.Println(names, len(names))

	// Slices

	var scores = []int{100, 50, 60}

	scores[2] = 25

	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

}
