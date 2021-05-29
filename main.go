// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that returns what you should do

package main

import (
	"fmt"
)

// This main function will calculate multuplication
func main() {
	// Defining variables
	var first int
	var second int
	var counter = 0
	var final = 0

	fmt.Println("Basic multipliction")
	fmt.Println()

	// User Input
	fmt.Print("First number: ")
	fmt.Scanln(&first)
	fmt.Println()

	fmt.Print("Second number: ")
	fmt.Scanln(&second)
	fmt.Println()

	// Does calculations and outputs result
	if second < 0 {
		for counter != second {
			final += first
			counter--
		}
		final = final - final - final
		fmt.Println("Multiplication: ", first, "x", second, "=", final)
	} else {
		for counter != second {
			final += first
			counter++
		}
		fmt.Println("Multiplication: ", first, "x", second, "=", final)
	}
}
