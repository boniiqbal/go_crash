package main

import "fmt"

func main() {
	// FruitArr := [2]string{"Apple", "Orange"}
	FruitArrSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(len(FruitArrSlice))
	fmt.Println(FruitArrSlice[0:2])
}
