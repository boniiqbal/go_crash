package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["bob"] = "bob@gmail.com"
	emails["tes"] = "tes@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["bob"])
}
