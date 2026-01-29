package main

import (
	"fmt"
	"os"
)

const Version = "1.0.0"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("Git Commit: test\n")
		fmt.Printf("Architecture: amd64\n")
		fmt.Printf("Go Version: go1.22.0\n")
		os.Exit(0)
	}

	fmt.Println("gstable - StableNet client")
	fmt.Println("Usage: gstable version")
}
