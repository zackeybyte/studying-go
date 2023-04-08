package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello go!"
	b := "hello"
	fmt.Println(strings.Contains(a, b))
}
