package files

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/glad-dev/sort-imports/sort"
)

// Format returns whether all files were formatted successfully.
func Format(m map[string]os.FileMode, moduleName string) bool {
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
			// We found a line that starts with "import ("
			start = i + 1
			for k := start; k < len(lines); k++ {
				l := strings.TrimSpace(lines[k])
				if len(l) == 0 {
					continue
				}

				if strings.HasPrefix(l, ")") {
					// We found the end of the import statement
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
		return fmt.Errorf("invalid start: %d", start)
	} else if end == 0 || end <= start {
		return fmt.Errorf("invalid end: %d", end)
	}

	// Sort imports
	sorted, err := sort.Imports(stmts, moduleName)
	if err != nil {
		return err
	}

	// Indent the import statements
	for i := range sorted {
		if len(strings.TrimSpace(sorted[i])) == 0 {
			// We don't want to indent new lines
			continue
		}

		sorted[i] = "\t" + sorted[i]
	}

	// Build new file content
	var newFile []string
	newFile = append(newFile, lines[:start]...) // Begin of the file
	newFile = append(newFile, sorted...)        // Import statements
	newFile = append(newFile, lines[end:]...)   // Rest of the file

	content := strings.Join(newFile, "\n")
	if string(f) == content {
		return nil
	}

	// Write back
	err = os.WriteFile(path, []byte(content), filePermissions)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}
