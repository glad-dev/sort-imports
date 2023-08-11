package sort

import (
	"fmt"
	"testing"
)

func TestIsFirstPartyLocalModule(t *testing.T) {
	moduleName := "moduleName"
	m := firstPartyModule(moduleName)

	for stmt, expected := range m {
		got := isFirstParty(stmt, moduleName)
		if got != expected {
			if expected {
				t.Errorf("'%s' should have been accepted, but was rejected", stmt)

				return
			}

			t.Errorf("'%s' should have been rejected, but was accepted", stmt)
		}
	}
}

func TestIsFirstPartyGlobalModule(t *testing.T) {
	// Module
	moduleName := "github.com/glad-dev/sort-imports"
	m := map[string]bool{
		// stdLib
		"\"testing\"":       false,
		"t \"testing\"":     false,
		"t2 \"testing\"":    false,
		"_ \"testing\"":     false,
		"\"path/filepath\"": false,
		// Own
		fmt.Sprintf("\"%s\"", moduleName):       true,
		fmt.Sprintf("\"%s/\"", moduleName):      true,
		fmt.Sprintf("\"%s/a\"", moduleName):     true,
		fmt.Sprintf("\"%s/a/\"", moduleName):    true,
		fmt.Sprintf("\"%s/a123/\"", moduleName): true,
		fmt.Sprintf("a \"%s/a\"", moduleName):   true,
		fmt.Sprintf("_ \"%s/a\"", moduleName):   true,
		// Third party
		"\"github.com\"":     false,
		"\"github.com/a\"":   false,
		"g \"github.com/a\"": false,
		"_ \"github.com/a\"": false,
	}

	for stmt, expected := range m {
		got := isFirstParty(stmt, moduleName)
		if got != expected {
			if expected {
				t.Errorf("%s should have been accepted, but was rejected", stmt)

				return
			}

			t.Errorf("%s should have been rejected, but was accepted", stmt)
		}
	}
}

func TestIsThirdParty(t *testing.T) {
	m := map[string]bool{
		// stdLib
		"\"test-ing\"":      false,
		"\"testing\"":       false,
		"t \"testing\"":     false,
		"t2 \"testing\"":    false,
		"_ \"testing\"":     false,
		"\"path/filepath\"": false,
		// Third party
		"\"github.com\"":      false,
		"\"github-com\"":      false,
		"\"gitlab.com/a\"":    true,
		"\"bitbucket.com/a\"": true,
		"\"github.com/a\"":    true,
		"g \"github.com/a\"":  true,
		"_ \"github.com/a\"":  true,
	}

	for stmt, expected := range m {
		got := isThirdParty(stmt)
		if got != expected {
			if expected {
				t.Errorf("%s should have been accepted, but was rejected", stmt)

				return
			}

			t.Errorf("'%s' should have been rejected, but was accepted", stmt)
		}
	}
}

func TestClassification(t *testing.T) {
	// Local module
	moduleName := "moduleName"
	m := firstPartyModule(moduleName)

	for stmt, expected := range m {
		n := determine(stmt, moduleName)
		old := isFirstParty(stmt, moduleName)

		if old != expected {
			t.Errorf("Old != expected!")

			return
		}

		if old && n != firstParty {
			t.Errorf("stmt '%s' differs; Old: %v, new: %v", stmt, old, n)
		} else if !old && n == firstParty {
			t.Errorf("stmt '%s' differs; Old: %v, new: %v", stmt, old, n)
		}
	}

	// Hosted module
	moduleName = "github.com/glad-dev/sort-imports"
	m = firstPartyModule(moduleName)

	for stmt, expected := range m {
		n := determine(stmt, moduleName)
		old := isFirstParty(stmt, moduleName)

		if old != expected {
			t.Errorf("Old != expected!")

			return
		}

		if old && n != firstParty {
			t.Errorf("stmt '%s' differs; Old: %v, new: %v", stmt, old, n)
		} else if !old && n == firstParty {
			t.Errorf("stmt '%s' differs; Old: %v, new: %v", stmt, old, n)
		}
	}

}

func firstPartyModule(moduleName string) map[string]bool {
	return map[string]bool{
		// stdLib
		"\"testing\"":       false,
		"t \"testing\"":     false,
		"t2 \"testing\"":    false,
		"_ \"testing\"":     false,
		"\"path/filepath\"": false,
		// Own
		fmt.Sprintf("\"%s\"", moduleName):       true,
		fmt.Sprintf("\"%s/\"", moduleName):      true,
		fmt.Sprintf("\"%s/a\"", moduleName):     true,
		fmt.Sprintf("\"%s/a/\"", moduleName):    true,
		fmt.Sprintf("\"%s/a123/\"", moduleName): true,
		fmt.Sprintf("a \"%s/a\"", moduleName):   true,
		fmt.Sprintf("_ \"%s/a\"", moduleName):   true,
		// Third party
		"\"github.com\"":     false,
		"\"github.com/a\"":   false,
		"g \"github.com/a\"": false,
		"_ \"github.com/a\"": false,
	}
}

func thirdPartyModule() map[string]bool {
	return map[string]bool{
		// stdLib
		"\"test-ing\"":      false,
		"\"testing\"":       false,
		"t \"testing\"":     false,
		"t2 \"testing\"":    false,
		"_ \"testing\"":     false,
		"\"path/filepath\"": false,
		// Third party
		"\"github.com\"":      false,
		"\"github-com\"":      false,
		"\"gitlab.com/a\"":    true,
		"\"bitbucket.com/a\"": true,
		"\"github.com/a\"":    true,
		"g \"github.com/a\"":  true,
		"_ \"github.com/a\"":  true,
	}
}

func stdLibList() map[string]party {
	std := []string{
		"archive/tar",
		"archive/zip",
		"bufio",
		"builtin",
		"bytes",
		"cmp",
		"compress/bzip2",
		"compress/flate",
		"compress/gzip",
		"compress/lzw",
		"compress/zlib",
		"container/heap",
		"container/list",
		"container/ring",
		"context",
		"crypto",
		"crypto/aes",
		"crypto/cipher",
		"crypto/des",
		"crypto/dsa",
		"crypto/ecdh",
		"crypto/ecdsa",
		"crypto/ed25519",
		"crypto/elliptic",
		"crypto/hmac",
		"crypto/md5",
		"crypto/rand",
		"crypto/rc4",
		"crypto/rsa",
		"crypto/sha1",
		"crypto/sha256",
		"crypto/sha512",
		"crypto/subtle",
		"crypto/tls",
		"crypto/x509",
		"crypto/x509/pkix",
		"database/sql",
		"database/sql/driver",
		"debug/buildinfo",
		"debug/dwarf",
		"debug/elf",
		"debug/gosym",
		"debug/macho",
		"debug/pe",
		"debug/plan9obj",
		"embed",
		"encoding",
		"encoding/ascii85",
		"encoding/asn1",
		"encoding/base32",
		"encoding/base64",
		"encoding/binary",
		"encoding/csv",
		"encoding/gob",
		"encoding/hex",
		"encoding/json",
		"encoding/pem",
		"encoding/xml",
		"errors",
		"expvar",
		"flag",
		"fmt",
		"go/ast",
		"go/build",
		"go/build/constant",
		"go/constant",
		"go/doc",
		"go/doc/comment",
		"go/format",
		"go/importer",
		"go/parser",
		"go/printer",
		"go/scanner",
		"go/token",
		"go/types",
		"hash",
		"hash/adler32",
		"hash/crc32",
		"hash/crc64",
		"hash/fnv",
		"hash/maphash",
		"html",
		"html/template",
		"image",
		"image/color",
		"image/color/palette",
		"image/draw",
		"image/gif",
		"image/jpeg",
		"image/png",
		"index/suffixarray",
		"io",
		"io/fs",
		"io/ioutil",
		"log",
		"log/slog",
		"log/syslog",
		"maps",
		"math",
		"math/big",
		"math/cmplx",
		"math/rand",
		"mime",
		"mime/multipart",
		"mime/quotedprintable",
		"net",
		"net/http",
		"net/http/cgi",
		"net/http/cookiejar",
		"net/http/fcgi",
		"net/http/httptest",
		"net/http/httptrace",
		"net/http/httputil",
		"net/http/pprof",
		"net/mail",
		"net/netip",
		"net/rpc",
		"net/rpc/jsonrpc",
		"net/smtp",
		"net/textproto",
		"net/url",
		"os",
		"os/exec",
		"os/signal",
		"os/user",
		"path",
		"path/filepath",
		"plugin",
		"reflect",
		"regexp",
		"regexp/syntax",
		"runtime",
		"runtime/cgo",
		"runtime/coverage",
		"runtime/debug",
		"runtime/metrics",
		"runtime/pprof",
		"runtime/race",
		"runtime/trace",
		"slices",
		"sort",
		"strconv",
		"strings",
		"sync",
		"sync/atomic",
		"syscall",
		"syscall/js",
		"testing",
		"testing/fstest",
		"testing/iotest",
		"testing/quick",
		"testing/slogtest",
		"text/scanner",
		"text/tabwriter",
		"text/template",
		"text/template/parse",
		"time",
		"time/tzdata",
		"unicode",
		"unicode/utf16",
		"unicode/utf8",
		"unsafe",
	}

	m := make(map[string]party, len(std))
	for _, s := range std {
		m[s] = stdLib
	}

	return m
}
