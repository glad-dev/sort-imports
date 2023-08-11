package sort

import "testing"

func TestTrimImport(t *testing.T) {
	valid := map[string]string{
		"\"\"":                        "",
		"\"test\"":                    "test",
		"\"t-e-s-t\"":                 "t-e-s-t",
		"someName \"moduleName\"":     "moduleName",
		"some_Name \"moduleName\"":    "moduleName",
		"some_Name123 \"moduleName\"": "moduleName",
	}

	invalid := []string{
		"",
		"\"",
		"\"Half quote with text",
		"\"asd\"a\"",
	}

	for stmt, expected := range valid {
		output, err := trimImport(stmt)
		if err != nil {
			t.Errorf("failed to trim import '%s': %s", stmt, err)
			continue
		}

		if output != expected {
			t.Errorf("expected: '%s', got '%s'", expected, output)
		}
	}

	for _, stmt := range invalid {
		output, err := trimImport(stmt)
		if err == nil {
			t.Errorf("Accepted invalid input '%s' and returned '%s'", stmt, output)
		}
	}
}
