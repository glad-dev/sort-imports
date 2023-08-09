package sort

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func Imports(imports []string, moduleName string) []string {
	std := make([]string, 0)
	first := make([]string, 0)
	third := make([]string, 0)

	for _, stmt := range imports {
		stmt = strings.TrimSpace(stmt)
		if len(stmt) == 0 {
			continue
		}

		if isFirstParty(stmt, moduleName) { // nolint:gocritic
			first = append(first, stmt)
		} else if isThirdParty(stmt) {
			third = append(third, stmt)
		} else {
			std = append(std, stmt)
		}
	}

	// Sort the imports
	sort.Strings(std)
	sort.Strings(first)
	sort.Strings(third)

	if len(std) > 0 && (len(first) > 0 || len(third) > 0) {
		std = append(std, "\n")
	}

	if len(first) > 0 && len(third) > 0 {
		first = append(first, "\n")
	}

	// Combine the slices into output array
	out := make([]string, 0)
	out = append(out, std...)
	out = append(out, first...)

	return append(out, third...)
}

type party int8

const (
	undetermined party = iota
	stdLib
	firstParty
	thirdParty
)

func determine(stmt string, moduleName string) party {
	stmt = trimImport(stmt)
	if len(stmt) == 0 {
		return undetermined
	}

	if len(moduleName) > 0 {
		// We need our module name to distinguish between first and third party
		if strings.HasPrefix(stmt, moduleName) {
			return firstParty
		}
	}

	b, _ := regexp.MatchString("^[a-zA-Z0-9]+\\.[a-zA-Z0-9]+/", stmt)
	if b {
		return thirdParty
	}

	return stdLib
}

func isFirstParty(stmt string, moduleName string) bool {
	if len(moduleName) == 0 {
		return false
	}

	stmt = trimImport(stmt)

	return strings.HasPrefix(stmt, "\""+moduleName)
}

func isThirdParty(stmt string) bool {
	stmt = trimImport(stmt)
	b, _ := regexp.MatchString("^\"[a-zA-Z0-9]+\\.[a-zA-Z0-9]+/", stmt)

	return b
}

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

	return stmt[:strings.Index(stmt, "\"")], nil
}
