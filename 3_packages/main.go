package main

import (
	"fmt"
	"math"

	res "github.com/boni/go_crash/3_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.1))
	fmt.Println(math.Sqrt(3))
	fmt.Println(res.Reverse("olleh"))
}
