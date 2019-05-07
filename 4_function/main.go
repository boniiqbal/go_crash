package main

import "fmt"

func greeting(name string) string {
	return "hello " + name
}

func sum(int1, int2 int) int {
	return int1 + int2
}

func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(greeting("wicak"))
}
