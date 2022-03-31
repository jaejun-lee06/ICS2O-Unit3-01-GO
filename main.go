// Created by: Jaejun Lee
// Created on: Mar 2022
//

package main

import (
	"fmt"
)

func main() {
	// This function does addition
	var aBase int
	var bBase int
	var height int

	// input
	fmt.Println("This program finds the area of a trapezoid")
	fmt.Println()
	fmt.Println("Enter the a base (cm): ")
	fmt.Scanln(&aBase)
	fmt.Println("Enter the b base (cm): ")
	fmt.Scanln(&bBase)
	fmt.Println("Enter the height (cm): ")
	fmt.Scanln(&height)

	// output
	fmt.Println("The area is: ", ((aBase+bBase)/2)*height, "cm²")
	fmt.Println("Done.")
}