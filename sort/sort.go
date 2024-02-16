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

		switch determine(stmt, moduleName) {
		case undetermined:
			// ToDo
			fallthrough

		case stdLib:
			std = append(std, stmt)

		case firstParty:
			first = append(first, stmt)

		case thirdParty:
			third = append(third, stmt)

		default:
			// Unreachable but return unsorted list, just in case
			return imports
		}
	}

	// Sort the imports
	sort.Strings(std)
	sort.Strings(first)
	sort.Strings(third)

	if len(std) > 0 && (len(first) > 0 || len(third) > 0) {
		std = append(std, "")
	}

	if len(first) > 0 && len(third) > 0 {
		first = append(first, "")
	}

	// Combine the slices into output array
	out := make([]string, 0)
	out = append(out, std...)
	out = append(out, first...)

	return append(out, third...)
}
