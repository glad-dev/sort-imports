package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/glad-dev/sort-imports/files"
	"github.com/glad-dev/sort-imports/module"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:     "sort-imports [path]",
		Short:   "Sorts import statements of Go files within the specified directory",
		Version: "0.1.0",
		Args:    cobra.ExactArgs(1),
		Run:     run,
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func run(_ *cobra.Command, args []string) {
	path, err := filepath.Abs(args[0])
	if err != nil {
		fmt.Printf("Failed to get the absolute path of '%s': %s", args[0], err)
		os.Exit(1)
	}

	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("The given path %s does not exist\n", path)
			os.Exit(1)
		}

		fmt.Printf("Failed to stat directory '%s': %s", path, err)
		os.Exit(1)
	}

	// Verify we were given a directory
	if !f.IsDir() {
		fmt.Printf("The provided path does not lead to a directory")
		os.Exit(1)
	}

	// Get module name
	moduleName, err := module.Name(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get all files with their permissions in the directory
	m := make(map[string]os.FileMode)
	m, err = files.Get(path, m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(m) == 0 {
		fmt.Printf("No go files found\n")
		return
	}

	if !files.Format(m, moduleName) {
		os.Exit(1)
	}

	if len(m) == 1 {
		fmt.Printf("Go file formated sucessfully.")
		return
	}

	fmt.Printf("All %d Go files were succesfully formatted.", len(m))
}
