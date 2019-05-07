package main

import "fmt"

func main() {
	x := 14
	y := 10

	if x < y {
		fmt.Printf("%d less than %d\n", x, y)
	} else {
		fmt.Printf("%d less than %d\n", y, x)
	}

	color := "red"

	if color == "yellow" {
		fmt.Println("Color is red")
	} else if color == "green" {
		fmt.Println("Color is green")
	} else {
		fmt.Println("Color has no color")
	}
}
