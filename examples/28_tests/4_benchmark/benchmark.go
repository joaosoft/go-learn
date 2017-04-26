// testes

package benchmark

import "testing"

// $ cd $GOPATH/src/gopl.io/ch11/word2
// $ go test -bench=.
// PASS
// BenchmarkIsPalindrome-8 1000000              1035 ns/op
// ok      gopl.io/ch11/word2      2.179s

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func IsPalindrome(teste string) {
	// do something...
}
