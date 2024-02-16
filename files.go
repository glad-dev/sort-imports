package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/glad-dev/sort-imports/sort"
)

// getFileList returns all files with a ".go" suffix in the given directory.
func getFileList(path string, m map[string]os.FileMode) (map[string]os.FileMode, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range dir {
		if entry.IsDir() {
			m, err = getFileList(filepath.Join(path, entry.Name()), m)
			if err != nil {
				return nil, err
			}

			continue
		}

		if !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		m[filepath.Join(path, entry.Name())] = info.Mode().Perm()
	}

	return m, nil
}

// formatFiles returns if all files were formatted successfully.
func formatFiles(m map[string]os.FileMode, moduleName string) bool {
	success := true
	wg := &sync.WaitGroup{}

	for path, mode := range m {
		wg.Add(1)

		go func(p string, fm os.FileMode) {
			defer wg.Done()

			err := handleFile(p, fm, moduleName)
			if err != nil {
				success = false
				fmt.Printf("Failed to format %s: %s\n", p, err)
			}
		}(path, mode)
	}

	wg.Wait()

	return success
}

func handleFile(path string, filePermissions os.FileMode, moduleName string) error {
	// Read file
	f, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}

	startRegex := regexp.MustCompile("^import\\s+(/\\*.*\\*/\\s*)?\\(") // Should match most import statement
	lines := strings.Split(string(f), "\n")
	stmts := make([]string, 0)
	start := 0
	end := 0

	for i, line := range lines {
		if startRegex.MatchString(line) {
			start = i + 1
			for k := i + 1; k < len(lines); k++ {
				l := strings.TrimSpace(lines[k])
				if len(l) == 0 {
					continue
				}

				if strings.HasPrefix(l, ")") {
					end = k

					break
				}

				stmts = append(stmts, l)
			}

			break
		}
	}

	if len(stmts) == 0 { // nolint: gocritic
		// File contains no multiline import statement
		return nil
	} else if start == 0 {
		return errors.New("invalid start") // ToDo
	} else if end == 0 || end <= start {
		return errors.New("invalid end") // ToDo
	}

	sorted := sort.Imports(stmts, moduleName)
	for i := range sorted {
		sorted[i] = "\t" + sorted[i]
	}

	// Build new file content
	newFile := lines[:start]
	newFile = append(newFile, sorted...)
	newFile = append(newFile, lines[end:]...)

	// Write files
	if false {
		err = os.WriteFile(path, []byte(strings.Join(newFile, "\n")), filePermissions)
		if err != nil {
			return fmt.Errorf("writing file: %w", err)
		}
	}

	return nil
}
