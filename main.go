package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	path, err := getPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	moduleName, err := getModuleName(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	m := make(map[string]os.FileMode)
	m, err = getFileList(path, m)
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

func getPath() (string, error) {
	var path string
	flag.StringVar(&path, "path", "", "path to format")
	flag.Parse()

	if len(path) == 0 {
		// No path given => Use working directory
		wd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("Failed to get working directory: %s\n", wd)
		}

		return wd, nil
	}

	// Path given => Check if it exists
	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("the path %s does not exist", path)
		}

		return "", err
	}

	// Verify we were given a directory
	if !f.IsDir() {
		return "", fmt.Errorf("the provided path does not lead to a directory")
	}

	return path, nil
}
