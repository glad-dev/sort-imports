package main

import (
	"fmt"
	"testing"
)

func TestIsFirstPartyLocalModule(t *testing.T) {
	// Module
	moduleName := "asd"
	m := map[string]bool{
		// stdLib
		"\"testing\"":       false,
		"t \"testing\"":     false,
		"t2 \"testing\"":    false,
		"_ \"testing\"":     false,
		"\"path/filepath\"": false,
		// Own
		fmt.Sprintf("\"%s\"", moduleName):       true,
		fmt.Sprintf("\"%s/\"", moduleName):      true,
		fmt.Sprintf("\"%s/a\"", moduleName):     true,
		fmt.Sprintf("\"%s/a/\"", moduleName):    true,
		fmt.Sprintf("\"%s/a123/\"", moduleName): true,
		fmt.Sprintf("a \"%s/a\"", moduleName):   true,
		fmt.Sprintf("_ \"%s/a\"", moduleName):   true,
		// Third party
		"\"github.com\"":     false,
		"\"github.com/a\"":   false,
		"g \"github.com/a\"": false,
		"_ \"github.com/a\"": false,
	}

	for stmt, expected := range m {
		got := isFirstParty(stmt, moduleName)
		if got != expected {
			if expected {
				t.Errorf("'%s' should have been accepted, but was rejected", stmt)
			}

			t.Errorf("'%s' should have been rejected, but was accepted", stmt)
		}
	}
}

func TestIsFirstPartyGlobalModule(t *testing.T) {
	// Module
	moduleName := "github.com/glad-dev/sort-imports"
	m := map[string]bool{
		// stdLib
		"\"testing\"":       false,
		"t \"testing\"":     false,
		"t2 \"testing\"":    false,
		"_ \"testing\"":     false,
		"\"path/filepath\"": false,
		// Own
		fmt.Sprintf("\"%s\"", moduleName):       true,
		fmt.Sprintf("\"%s/\"", moduleName):      true,
		fmt.Sprintf("\"%s/a\"", moduleName):     true,
		fmt.Sprintf("\"%s/a/\"", moduleName):    true,
		fmt.Sprintf("\"%s/a123/\"", moduleName): true,
		fmt.Sprintf("a \"%s/a\"", moduleName):   true,
		fmt.Sprintf("_ \"%s/a\"", moduleName):   true,
		// Third party
		"\"github.com\"":     false,
		"\"github.com/a\"":   false,
		"g \"github.com/a\"": false,
		"_ \"github.com/a\"": false,
	}

	for stmt, expected := range m {
		got := isFirstParty(stmt, moduleName)
		if got != expected {
			if expected {
				t.Errorf("%s should have been accepted, but was rejected", stmt)
			}

			t.Errorf("%s should have been rejected, but was accepted", stmt)
		}
	}
}

func TestIsThirdParty(t *testing.T) {
	m := map[string]bool{
		// stdLib
		"\"test-ing\"":      false,
		"\"testing\"":       false,
		"t \"testing\"":     false,
		"t2 \"testing\"":    false,
		"_ \"testing\"":     false,
		"\"path/filepath\"": false,
		// Third party
		"\"github.com\"":      false,
		"\"github-com\"":      false,
		"\"gitlab.com/a\"":    true,
		"\"bitbucket.com/a\"": true,
		"\"github.com/a\"":    true,
		"g \"github.com/a\"":  true,
		"_ \"github.com/a\"":  true,
	}

	for stmt, expected := range m {
		got := isThirdParty(stmt)
		if got != expected {
			if expected {
				t.Errorf("%s should have been accepted, but was rejected", stmt)
			}

			t.Errorf("'%s' should have been rejected, but was accepted", stmt)
		}
	}
}

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
				"\n",
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
				"\n",
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
				"\n",
				"\"github.com/glad-dev/sort-imports/\"",
				"\n",
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
				"\n",
				"\"github.com/glad-dev/sort-imports/\"",
				"\n",
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
	}

	for _, c := range testCases {
		sorted := sortImports(c.imports, c.moduleName)
		if !compareStringArray(sorted, c.expected) {
			t.Errorf("Expected (%d): %s, got (%d): %v", len(c.expected), str(c.expected), len(sorted), str(sorted))
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

func str(a []string) string {
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
