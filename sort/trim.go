package sort

import (
	"fmt"
	"strings"
)

func trimImport(stmt string) string {
	// Remove custom naming
	i := strings.Index(stmt, "\"")
	if i == -1 {
		return stmt
	}

	return strings.TrimSpace(stmt[i:])
}

func trimI(stmt string) (string, error) {
	if strings.Count(stmt, "\"") != 2 {
		return "", fmt.Errorf("import statement '%s' is malformed: does not contain two \"", stmt)
	}

	// Remove import name and first "
	index := strings.Index(stmt, "\"")
	stmt = stmt[index+1:] // index+1 is in range of stmt since we have at least two " and index returns the first occurrence

	// We know that strings.Index returns a valid position since we know that the input string contains two "
	return stmt[:strings.Index(stmt, "\"")], nil // nolint:gocritic
}
