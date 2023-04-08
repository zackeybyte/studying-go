package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello go!"
	fmt.Println(strings.ReplaceAll(a, "l", " "))
}
