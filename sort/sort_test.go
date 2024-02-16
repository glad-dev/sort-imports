package sort

import (
	"fmt"
	"testing"
)

func TestSortImports(t *testing.T) {
	type testCase struct {
		moduleName string
		imports    []string
		expected   []string
	}

	testCases := []testCase{
		{
			moduleName: "",
			imports: []string{
				"\"testing\"",
				"\"fmt\"",
			},
			expected: []string{
				"\"fmt\"",
				"\"testing\"",
			},
		},
		{
			moduleName: "",
			imports: []string{
				"\"fmt\"",
				"",
				"\"testing\"",
			},
			expected: []string{
				"\"fmt\"",
				"\"testing\"",
			},
		},
		{
			moduleName: "",
			imports: []string{
				"\"testing\"",
				"\"fmt\"",
			},
			expected: []string{
				"\"fmt\"",
				"\"testing\"",
			},
		},
		{
			moduleName: "",
			imports: []string{
				"\"testing\"",
				"",
				"\"fmt\"",
			},
			expected: []string{
				"\"fmt\"",
				"\"testing\"",
			},
		},
		{
			moduleName: "",
			imports: []string{
				"\"fmt\"",
				"\"github.com/glad-dev/sort-imports/\"",
			},
			expected: []string{
				"\"fmt\"",
				"",
				"\"github.com/glad-dev/sort-imports/\"",
			},
		},
		{
			moduleName: "github.com/glad-dev/sort-imports",
			imports: []string{
				"\"fmt\"",
				"\"github.com/glad-dev/sort-imports/\"",
			},
			expected: []string{
				"\"fmt\"",
				"",
				"\"github.com/glad-dev/sort-imports/\"",
			},
		},
		{
			moduleName: "github.com/glad-dev/sort-imports",
			imports: []string{
				"\"fmt\"",
				"\"github.com/glad-dev/other-repo/\"",
				"\"testing\"",
				"\"github.com/glad-dev/sort-imports/\"",
			},
			expected: []string{
				"\"fmt\"",
				"\"testing\"",
				"",
				"\"github.com/glad-dev/sort-imports/\"",
				"",
				"\"github.com/glad-dev/other-repo/\"",
			},
		},
		{
			moduleName: "github.com/glad-dev/sort-imports",
			imports: []string{
				"\"fmt\"",
				"\"github.com/glad-dev/other-repo/\"",
				"\"testing\"",
				"\"github.com/glad-dev/sort-imports/\"",
			},
			expected: []string{
				"\"fmt\"",
				"\"testing\"",
				"",
				"\"github.com/glad-dev/sort-imports/\"",
				"",
				"\"github.com/glad-dev/other-repo/\"",
			},
		},
		{
			moduleName: "github.com/glad-dev/sort-imports",
			imports: []string{
				"\"errors\"",
				"\"fmt\"",
				"\"os\"",
				"\"path/filepath\"",
				"\"strings\"",
				"\"sync\"",
			},
			expected: []string{
				"\"errors\"",
				"\"fmt\"",
				"\"os\"",
				"\"path/filepath\"",
				"\"strings\"",
				"\"sync\"",
			},
		},
		{
			moduleName: "github.com/glad-dev/sort-imports",
			imports: []string{
				"\"github.com/glad-dev/sort-imports/sub2\"",
				"x \"github.com/glad-dev/sort-imports/sub1\"",
			},
			expected: []string{
				"x \"github.com/glad-dev/sort-imports/sub1\"",
				"\"github.com/glad-dev/sort-imports/sub2\"",
			},
		},
	}

	for _, c := range testCases {
		sorted, err := Imports(c.imports, c.moduleName)
		if err != nil {
			t.Errorf("Function returned an error: %s", err)
		}

		if !compareStringArray(sorted, c.expected) {
			t.Errorf("Expected (%d): %s, got (%d): %v", len(c.expected), replaceSpecial(c.expected), len(sorted), replaceSpecial(sorted))
		}
	}
}

func compareStringArray(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func replaceSpecial(a []string) string {
	for i, s := range a {
		switch s {
		case "\n":
			a[i] = "\"\\n\""
		case "\t":
			a[i] = "\"\\t\""
		}
	}

	return fmt.Sprintf("%v", a)
}
