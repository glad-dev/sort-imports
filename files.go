package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

func do(path string, moduleName string) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("reading dir: %w", err)
	}

	errChan := make(chan error, 1)
	wg := &sync.WaitGroup{}

	for _, entry := range dir {
		if entry.IsDir() {
			_ = do(filepath.Join(path, entry.Name()), moduleName) // ToDo
		}

		if strings.HasSuffix(entry.Name(), ".go") {
			fmt.Printf("Found go file: %s\n", entry.Name())
			go func() {
				wg.Add(1)
				handleFile(filepath.Join(path, entry.Name()), moduleName, errChan)
				wg.Done()
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

func handleFile(path string, moduleName string, c chan error) {
	// Read file
	f, err := os.ReadFile(path)
	if err != nil {
		c <- fmt.Errorf("opening file: %s", err)
		return
	}

	re := regexp.MustCompile("import\\s*\\(\\s*\"[^\"]+\"(?:\\s+\"[^\"]+\")*\\s*\\)\n")
	r := re.Find(f)
	if r == nil {
		fmt.Printf("No import in %s\n", path)
		return
	}

	lines := strings.Split(string(f), "\n")
	stmts := make([]string, 0)
	for i, line := range lines {
		if strings.Contains(line, "import (") {
			for k := i + 1; k < len(lines); k++ {
				l := strings.TrimSpace(lines[k])

				if l == ")" {
					break
				}

				stmts = append(stmts, l)
			}
		}
	}

	s := strings.Split(string(r), "\n")
	if len(s) < 2 {
		c <- fmt.Errorf("invalid import statement: %s", r)
		return
	}

	// Remove first line: "import (" and last line: ")"
	s = s[1:]

	fmt.Printf("Stmts: %v\nRegex-%d: %v\n", stmts, len(s), str(s))
}

func str(a []string) string {
	for i, s := range a {
		switch s {
		case "\n":
			a[i] = fmt.Sprintf("%d - \"\\n\"", i)
		case "\t":
			a[i] = fmt.Sprintf("%d - \"\\t\"", i)
		default:
			
		}
	}

	return fmt.Sprintf("%v", a)
}
