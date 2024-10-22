package main

import (
	"fmt"
	"log"
	"os"

	"linear-stats/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("\x1b[33mUsage: go run . data.txt\033[0m")
		return
	}

	// Read data.txt
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	functions.Data_Manipulation(string(file))
}
