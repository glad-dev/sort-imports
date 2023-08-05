package main

import (
	"regexp"
	"sort"
	"strings"
)

func s(imports []string, moduleName string) []string {
	stdLib := make([]string, 0)
	firstParty := make([]string, 0)
	thirdParty := make([]string, 0)

	for _, s := range imports {
		if isFirstParty(s, moduleName) {
			firstParty = append(firstParty, s)
		} else if isThirdParty(s) {
			thirdParty = append(thirdParty, s)
		} else {
			stdLib = append(stdLib, s)
		}
	}

	// Sort the imports
	sort.Strings(stdLib)
	sort.Strings(firstParty)
	sort.Strings(thirdParty)

	if len(stdLib) > 0 {
		stdLib = append(stdLib, "\n")
	}

	if len(firstParty) > 0 {
		firstParty = append(firstParty, "\n")
	}

	// Combine the slices into output array
	out := make([]string, 0)
	out = append(out, stdLib...)
	out = append(out, firstParty...)

	return append(out, thirdParty...)
}

func isFirstParty(stmt string, moduleName string) bool {
	stmt = removeCustomNaming(stmt)
	return strings.HasPrefix(stmt, moduleName)
}

func isThirdParty(stmt string) bool {
	b, _ := regexp.MatchString("^[a-zA-Z]+.[a-zA-Z]/.+", stmt)

	return b
}

func removeCustomNaming(stmt string) string {
	// Remove custom naming
	i := strings.Index(stmt, "\"")
	if i == -1 {
		panic("No \"") // ToDo: Error msg
	} else if len(stmt) < i+1 {
		panic("Too short")
	}

	return stmt[i+1:]
}
