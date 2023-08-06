package main

import (
	"regexp"
	"sort"
	"strings"
)

func sortImports(imports []string, moduleName string) []string {
	stdLib := make([]string, 0)
	firstParty := make([]string, 0)
	thirdParty := make([]string, 0)

	for _, stmt := range imports {
		stmt = strings.TrimSpace(stmt)
		if len(stmt) == 0 {
			continue
		}

		if isFirstParty(stmt, moduleName) { // nolint:gocritic
			firstParty = append(firstParty, stmt)
		} else if isThirdParty(stmt) {
			thirdParty = append(thirdParty, stmt)
		} else {
			stdLib = append(stdLib, stmt)
		}
	}

	// Sort the imports
	sort.Strings(stdLib)
	sort.Strings(firstParty)
	sort.Strings(thirdParty)

	if len(stdLib) > 0 && (len(firstParty) > 0 || len(thirdParty) > 0) {
		stdLib = append(stdLib, "\n")
	}

	if len(firstParty) > 0 && len(thirdParty) > 0 {
		firstParty = append(firstParty, "\n")
	}

	// Combine the slices into output array
	out := make([]string, 0)
	out = append(out, stdLib...)
	out = append(out, firstParty...)

	return append(out, thirdParty...)
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

	return stmt[i:]
}
