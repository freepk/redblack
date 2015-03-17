package redblack

import (
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	n := newNode(0, black)
	for i := 0; i <= 10000; i++ {
		n, _ = n.insert(rand.Int())
	}
	n, _ = n.insert(0)
	n, _ = n.insert(0)
}

func BenchmarkInsSeq(b *testing.B) {
	n := newNode(0, black)
	for i := 0; i < b.N; i++ {
		n, _ = n.insert(i)
	}
}

func BenchmarkInsRnd(b *testing.B) {
	f := false
	c := 0
	n := newNode(0, black)
	for i := 0; i < b.N; i++ {
		if n, f = n.insert(rand.Int()); f == true {
			c++
		}
	}
	b.Logf("b.N: %d, c: %d", b.N, c)
}
