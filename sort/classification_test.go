package sort

import (
	"fmt"
	"testing"
)

func TestStdLibClassification(t *testing.T) {
	std := stdLibList()
	for stmt, expected := range std {
		noName := determine(stmt, "")
		localName := determine(stmt, "sort-imports")
		moduleName := determine(stmt, "github.com/sort-imports")
		sameName := determine(stmt, stmt[1:len(stmt)-1]) // Check if shadowing a package name works

		if noName != expected {
			t.Errorf("stmt '%s' is not classified as stdLib: %s", stmt, noName.String())
		}

		if localName != expected {
			t.Errorf("stmt '%s' is not classified as stdLib: %s", stmt, localName.String())
		}

		if moduleName != expected {
			t.Errorf("stmt '%s' is not classified as stdLib: %s", stmt, moduleName.String())
		}

		if sameName != firstParty {
			t.Errorf("Error with shadowing: '%s' is not classified as first party: %s", stmt, moduleName.String())
		}
	}
}

func TestFirstPartyClassification(t *testing.T) {
	moduleName := "sort-imports"
	for stmt, expected := range firstAndThirdParty(moduleName) {
		if d := determine(stmt, moduleName); d != expected {
			t.Errorf("stmt '%s' is incorrectly classified as %s instead of %s", stmt, d.String(), expected.String())
		}
	}

	moduleName = "github.com/glad-dev/sort-imports"
	for stmt, expected := range firstAndThirdParty(moduleName) {
		if d := determine(stmt, moduleName); d != expected {
			t.Errorf("stmt '%s' is incorrectly classified as %s instead of %s", stmt, d.String(), expected.String())
		}
	}
}

func firstAndThirdParty(moduleName string) map[string]party {
	// We don't need to test first party packages here since they have a separate test.
	return map[string]party{
		fmt.Sprintf("\"%s\"", moduleName):       firstParty,
		fmt.Sprintf("\"%s/\"", moduleName):      firstParty,
		fmt.Sprintf("\"%s/a\"", moduleName):     firstParty,
		fmt.Sprintf("\"%s/a/\"", moduleName):    firstParty,
		fmt.Sprintf("\"%s/a123/\"", moduleName): firstParty,
		fmt.Sprintf("a \"%s/a\"", moduleName):   firstParty,
		fmt.Sprintf("_ \"%s/a\"", moduleName):   firstParty,

		"\"gitlab.com/a\"":    thirdParty,
		"\"bitbucket.com/a\"": thirdParty,
		"\"github.com/a\"":    thirdParty,
	}
}

func stdLibList() map[string]party {
	return map[string]party{
		"\"archive/tar\"":          stdLib,
		"\"archive/zip\"":          stdLib,
		"\"bufio\"":                stdLib,
		"\"builtin\"":              stdLib,
		"\"bytes\"":                stdLib,
		"\"cmp\"":                  stdLib,
		"\"compress/bzip2\"":       stdLib,
		"\"compress/flate\"":       stdLib,
		"\"compress/gzip\"":        stdLib,
		"\"compress/lzw\"":         stdLib,
		"\"compress/zlib\"":        stdLib,
		"\"container/heap\"":       stdLib,
		"\"container/list\"":       stdLib,
		"\"container/ring\"":       stdLib,
		"\"context\"":              stdLib,
		"\"crypto\"":               stdLib,
		"\"crypto/aes\"":           stdLib,
		"\"crypto/cipher\"":        stdLib,
		"\"crypto/des\"":           stdLib,
		"\"crypto/dsa\"":           stdLib,
		"\"crypto/ecdh\"":          stdLib,
		"\"crypto/ecdsa\"":         stdLib,
		"\"crypto/ed25519\"":       stdLib,
		"\"crypto/elliptic\"":      stdLib,
		"\"crypto/hmac\"":          stdLib,
		"\"crypto/md5\"":           stdLib,
		"\"crypto/rand\"":          stdLib,
		"\"crypto/rc4\"":           stdLib,
		"\"crypto/rsa\"":           stdLib,
		"\"crypto/sha1\"":          stdLib,
		"\"crypto/sha256\"":        stdLib,
		"\"crypto/sha512\"":        stdLib,
		"\"crypto/subtle\"":        stdLib,
		"\"crypto/tls\"":           stdLib,
		"\"crypto/x509\"":          stdLib,
		"\"crypto/x509/pkix\"":     stdLib,
		"\"database/sql\"":         stdLib,
		"\"database/sql/driver\"":  stdLib,
		"\"debug/buildinfo\"":      stdLib,
		"\"debug/dwarf\"":          stdLib,
		"\"debug/elf\"":            stdLib,
		"\"debug/gosym\"":          stdLib,
		"\"debug/macho\"":          stdLib,
		"\"debug/pe\"":             stdLib,
		"\"debug/plan9obj\"":       stdLib,
		"\"embed\"":                stdLib,
		"\"encoding\"":             stdLib,
		"\"encoding/ascii85\"":     stdLib,
		"\"encoding/asn1\"":        stdLib,
		"\"encoding/base32\"":      stdLib,
		"\"encoding/base64\"":      stdLib,
		"\"encoding/binary\"":      stdLib,
		"\"encoding/csv\"":         stdLib,
		"\"encoding/gob\"":         stdLib,
		"\"encoding/hex\"":         stdLib,
		"\"encoding/json\"":        stdLib,
		"\"encoding/pem\"":         stdLib,
		"\"encoding/xml\"":         stdLib,
		"\"errors\"":               stdLib,
		"\"expvar\"":               stdLib,
		"\"flag\"":                 stdLib,
		"\"fmt\"":                  stdLib,
		"\"go/ast\"":               stdLib,
		"\"go/build\"":             stdLib,
		"\"go/build/constant\"":    stdLib,
		"\"go/constant\"":          stdLib,
		"\"go/doc\"":               stdLib,
		"\"go/doc/comment\"":       stdLib,
		"\"go/format\"":            stdLib,
		"\"go/importer\"":          stdLib,
		"\"go/parser\"":            stdLib,
		"\"go/printer\"":           stdLib,
		"\"go/scanner\"":           stdLib,
		"\"go/token\"":             stdLib,
		"\"go/types\"":             stdLib,
		"\"hash\"":                 stdLib,
		"\"hash/adler32\"":         stdLib,
		"\"hash/crc32\"":           stdLib,
		"\"hash/crc64\"":           stdLib,
		"\"hash/fnv\"":             stdLib,
		"\"hash/maphash\"":         stdLib,
		"\"html\"":                 stdLib,
		"\"html/template\"":        stdLib,
		"\"image\"":                stdLib,
		"\"image/color\"":          stdLib,
		"\"image/color/palette\"":  stdLib,
		"\"image/draw\"":           stdLib,
		"\"image/gif\"":            stdLib,
		"\"image/jpeg\"":           stdLib,
		"\"image/png\"":            stdLib,
		"\"index/suffixarray\"":    stdLib,
		"\"io\"":                   stdLib,
		"\"io/fs\"":                stdLib,
		"\"io/ioutil\"":            stdLib,
		"\"log\"":                  stdLib,
		"\"log/slog\"":             stdLib,
		"\"log/syslog\"":           stdLib,
		"\"maps\"":                 stdLib,
		"\"math\"":                 stdLib,
		"\"math/big\"":             stdLib,
		"\"math/cmplx\"":           stdLib,
		"\"math/rand\"":            stdLib,
		"\"mime\"":                 stdLib,
		"\"mime/multipart\"":       stdLib,
		"\"mime/quotedprintable\"": stdLib,
		"\"net\"":                  stdLib,
		"\"net/http\"":             stdLib,
		"\"net/http/cgi\"":         stdLib,
		"\"net/http/cookiejar\"":   stdLib,
		"\"net/http/fcgi\"":        stdLib,
		"\"net/http/httptest\"":    stdLib,
		"\"net/http/httptrace\"":   stdLib,
		"\"net/http/httputil\"":    stdLib,
		"\"net/http/pprof\"":       stdLib,
		"\"net/mail\"":             stdLib,
		"\"net/netip\"":            stdLib,
		"\"net/rpc\"":              stdLib,
		"\"net/rpc/jsonrpc\"":      stdLib,
		"\"net/smtp\"":             stdLib,
		"\"net/textproto\"":        stdLib,
		"\"net/url\"":              stdLib,
		"\"os\"":                   stdLib,
		"\"os/exec\"":              stdLib,
		"\"os/signal\"":            stdLib,
		"\"os/user\"":              stdLib,
		"\"path\"":                 stdLib,
		"\"path/filepath\"":        stdLib,
		"\"plugin\"":               stdLib,
		"\"reflect\"":              stdLib,
		"\"regexp\"":               stdLib,
		"\"regexp/syntax\"":        stdLib,
		"\"runtime\"":              stdLib,
		"\"runtime/cgo\"":          stdLib,
		"\"runtime/coverage\"":     stdLib,
		"\"runtime/debug\"":        stdLib,
		"\"runtime/metrics\"":      stdLib,
		"\"runtime/pprof\"":        stdLib,
		"\"runtime/race\"":         stdLib,
		"\"runtime/trace\"":        stdLib,
		"\"slices\"":               stdLib,
		"\"sort\"":                 stdLib,
		"\"strconv\"":              stdLib,
		"\"strings\"":              stdLib,
		"\"sync\"":                 stdLib,
		"\"sync/atomic\"":          stdLib,
		"\"syscall\"":              stdLib,
		"\"syscall/js\"":           stdLib,
		"\"testing\"":              stdLib,
		"\"testing/fstest\"":       stdLib,
		"\"testing/iotest\"":       stdLib,
		"\"testing/quick\"":        stdLib,
		"\"testing/slogtest\"":     stdLib,
		"\"text/scanner\"":         stdLib,
		"\"text/tabwriter\"":       stdLib,
		"\"text/template\"":        stdLib,
		"\"text/template/parse\"":  stdLib,
		"\"time\"":                 stdLib,
		"\"time/tzdata\"":          stdLib,
		"\"unicode\"":              stdLib,
		"\"unicode/utf16\"":        stdLib,
		"\"unicode/utf8\"":         stdLib,
		"\"unsafe\"":               stdLib,
	}
}
