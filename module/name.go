package module

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Name(path string) (string, error) {
	d, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	for _, entry := range d {
		if entry.Name() == "go.mod" {
			return readModFile(filepath.Join(path, entry.Name()))
		}
	}

	parent := filepath.Dir(path)
	if len(parent) == 0 || parent == path {
		return "", fmt.Errorf("no go.mod file was found")
	}

	return Name(parent)
}

func readModFile(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading go.mod: %w", err)
	}

	lines := strings.Split(string(f), "\n")
	if len(lines) == 0 {
		return "", fmt.Errorf("go.mod is empty")
	}

	location := -1
	for i, line := range lines {
		if strings.HasPrefix(line, "module") {
			location = i
			break
		}
	}

	if location == -1 {
		return "", fmt.Errorf("go.mod contains no 'module'")
	}

	name := strings.TrimSpace(lines[location][len("module"):])
	if len(name) == 0 {
		return "", fmt.Errorf("go.mod contains empty module name")
	}

	return name, nil
}
