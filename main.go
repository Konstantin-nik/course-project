package main

import (
	"fmt"
)

func main() {
	fmt.Print("Welcome to the arena")
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
