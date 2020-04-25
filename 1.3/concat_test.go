// Exercise 1.3 compares string concatenation techniques
//
// Run with go test -bench=.
package concat_test

import (
	"strings"
	"testing"
)

var args = []string{"what", "do", "these", "arguments", "do", "help", "please"}

func concat(args []string) {
	r, sep := "", ""
	for _, a := range args {
		r += sep + a
		sep = " "
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}
