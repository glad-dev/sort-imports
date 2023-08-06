package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestFindModFile(t *testing.T) {
	moduleName := "github.com/glad-dev/sort-imports"

	run := func(path string) {
		name, err := getModuleName(path)
		if err != nil {
			t.Errorf("Failed to get module name: %s", err)
		} else if name != moduleName {
			t.Errorf("Module name mismatch; Expected: %s, got %s", moduleName, name)
		}
	}

	dir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Errorf("Failed to create tmp dir: %s", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Printf("Failed to delete tmp dir: %s\n", err)
		}
	}(dir)

	// Directory contains no go.mod file => should fail
	name, err := getModuleName(dir)
	if err == nil {
		t.Errorf("Found a go.mod file in an empty directory. Module name: %s", name)
	}

	f, err := os.OpenFile(filepath.Join(dir, "go.mod"), os.O_CREATE|os.O_WRONLY, 0o664)
	if err != nil {
		t.Errorf("Faild to create go.mod: %s", err)
	}

	_, err = f.WriteString(fmt.Sprintf("module %s\ngo 1.20\n", moduleName))
	if err != nil {
		t.Errorf("Faild to write to go.mod: %s", err)
	}

	run(dir)

	// Create sub directories
	sub, err := os.MkdirTemp(dir, "")
	if err != nil {
		t.Errorf("Failed to create sub-directory: %s", err)
	}

	run(sub)

	sub2, err := os.MkdirTemp(sub, "")
	if err != nil {
		t.Errorf("Failed to create sub-sub-directory: %s", err)
	}

	run(sub2)
}
