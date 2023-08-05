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
		"\"testing\"":    false,
		"t \"testing\"":  false,
		"t2 \"testing\"": false,
		"_ \"testing\"":  false,
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
		"\"testing\"":    false,
		"t \"testing\"":  false,
		"t2 \"testing\"": false,
		"_ \"testing\"":  false,
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
		"\"test-ing\"":   false,
		"\"testing\"":    false,
		"t \"testing\"":  false,
		"t2 \"testing\"": false,
		"_ \"testing\"":  false,
		// Third party
		"\"github.com\"":     false,
		"\"github-com\"":     false,
		"\"github.com/a\"":   true,
		"g \"github.com/a\"": true,
		"_ \"github.com/a\"": true,
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
