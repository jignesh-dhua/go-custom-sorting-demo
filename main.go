package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("main started")

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	fmt.Println("main ended")
}
