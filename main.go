package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get working directory: %s\n", wd)
		os.Exit(1)
	}

	f, err := getModuleName(wd)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	err = do(wd, f)
	if err != nil {
		fmt.Printf("Failed to do work: %s\n", err)
		os.Exit(1)
	}
}
