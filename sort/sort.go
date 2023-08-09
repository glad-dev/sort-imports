package sort

import (
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
