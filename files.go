package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/glad-dev/sort-imports/sort"
)

func do(path string, moduleName string) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("reading dir: %w", err)
	}

	errChan := make(chan error, 100)
	wg := &sync.WaitGroup{}

	for _, entry := range dir {
		if entry.IsDir() {
			_ = do(filepath.Join(path, entry.Name()), moduleName) // ToDo
		}

		info, err := entry.Info()
		if err != nil {
			// Can't get file info => We can't get file permissions => Abort
			continue // ToDo: Don't ignore error
		}

		if strings.HasSuffix(entry.Name(), ".go") {
			func() {
				wg.Add(1)
				defer wg.Done()

				err := handleFile(filepath.Join(path, entry.Name()), moduleName, info.Mode().Perm())
				if err != nil {
					errChan <- err
				}
			}()
		}
	}

	wg.Wait()

	select {
	case e := <-errChan:
		return e

	default:
		return nil
	}
}

func handleFile(path string, moduleName string, filePermissions os.FileMode) error {
	// Read file
	f, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}

	lines := strings.Split(string(f), "\n")
	stmts := make([]string, 0)
	start := 0
	end := 0

	for i, line := range lines {
		if strings.Contains(line, "import (") {
			start = i + 1
			for k := i + 1; k < len(lines); k++ {
				l := strings.TrimSpace(lines[k])

				if l == ")" {
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
