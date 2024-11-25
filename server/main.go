package main

import (
	"fmt"

	firecharge "github.com/telecter/firecharge/internal"
)

func main() {
	fmt.Println("*firecharge*")
	fmt.Println("Starting on port 3000...")
	firecharge.Start()
}
