package sort

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func Imports(imports []string, moduleName string) ([]string, error) {
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
			return imports, errors.New("unreachable")
		}
	}

	// Sort the imports

	if err := sortImports(std); err != nil {
		return nil, err
	}

	if err := sortImports(first); err != nil {
		return nil, err
	}

	if err := sortImports(third); err != nil {
		return nil, err
	}

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

	return append(out, third...), nil
}

func sortImports(imports []string) error {
	r := regexp.MustCompile("^.+ \"") // Matches all named imports
	dec := make(map[string]string)    // url => name
	for i, l := range imports {
		if r.MatchString(l) {
			split := strings.SplitN(l, " ", 2)
			if len(split) != 2 {
				return fmt.Errorf("sortImports: split failed for %s", l)
			}

			split[0] = strings.TrimSpace(split[0])
			split[1] = strings.TrimSpace(split[1])

			imports[i] = split[1]
			dec[split[1]] = split[0]
		}
	}

	sort.Strings(imports)
	if len(dec) == 0 {
		// No named imports
		return nil
	}

	for i, l := range imports {
		name, ok := dec[l]
		if !ok {
			continue
		}

		imports[i] = fmt.Sprintf("%s %s", name, l)
		delete(dec, l)
	}

	// Sanity check to verify that all names have been re-added
	if len(dec) != 0 {
		return fmt.Errorf("sortImports: didn't re-add all names %+v", dec)
	}

	return nil
}
