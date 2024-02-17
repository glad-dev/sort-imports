package files

import (
	"os"
	"path/filepath"
	"strings"
)

// Get returns all files with a ".go" suffix in the given directory.
func Get(path string, m map[string]os.FileMode) (map[string]os.FileMode, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range dir {
		if entry.IsDir() {
			m, err = Get(filepath.Join(path, entry.Name()), m)
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
