package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	m1 := 2
// 	m2 := 3
// 	fmt.Println(m1 + m2)
// }

func main() {
	m1 := "My name"
	m1 = "name"
	fmt.Println(strings.ReplaceAll(m1, "m", "NO"))
}
