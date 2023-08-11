package sort

import (
	"regexp"
	"strings"
)

type party struct {
	slug string
}

func (p *party) String() string {
	return p.slug
}

var (
	undetermined = party{slug: "undetermined"}
	stdLib       = party{slug: "stdLib"}
	firstParty   = party{slug: "firstParty"}
	thirdParty   = party{slug: "thirdParty"}
)

func determine(stmt string, moduleName string) party {
	stmt, err := trimI(stmt)
	if err != nil || len(stmt) == 0 {
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
