package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
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

		info, _ := entry.Info()
		s := info.Mode().Perm().String()
		fmt.Printf("Mode: %s\n", s)

		if strings.HasSuffix(entry.Name(), ".go") {
			func() {
				wg.Add(1)
				defer wg.Done()

				err := handleFile(filepath.Join(path, entry.Name()), moduleName)
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

func handleFile(path string, moduleName string) error {
	// Read file
	f, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("opening file: %s", err)
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

	if len(stmts) == 0 {
		// No multiline import statement
		return nil
	} else if start == 0 {
		return errors.New("invalid start") // ToDo
	} else if end == 0 || end <= start {
		return errors.New("invalid end") // ToDo
	}

	sorted := sortImports(stmts, moduleName)
	for i, _ := range sorted {
		sorted[i] = "\t" + sorted[i]
	}

	// Build new file content
	newFile := lines[:start]
	newFile = append(newFile, sorted...)
	newFile = append(newFile, lines[end:]...)

	// Write files
	//os.WriteFile(path, []byte(strings.Join(newFile, "\n")), 0o666)

	return nil
}
