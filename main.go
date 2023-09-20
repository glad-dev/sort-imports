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

	moduleName, err := getModuleName(wd)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	m := make(map[string]os.FileMode)
	m, err = getFileList(wd, m)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if len(m) == 0 {
		fmt.Printf("No go files found\n")

		return
	}

	if !formatFiles(m, moduleName) {
		os.Exit(1)
	}

	if len(m) == 1 {
		fmt.Printf("Go file formated sucessfully.")

		return
	}

	fmt.Printf("All %d Go files were succesfully formatted.", len(m))
}
