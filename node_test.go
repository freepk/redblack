package redblack

import (
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	n := newNode(10, black)
	for i := 20; i <= 100; i += 10 {
		n, _ = n.insert(i)
	}
}

func BenchmarkInsSeq(b *testing.B) {
	n := newNode(0, black)
	for i := 0; i < b.N; i++ {
		n, _ = n.insert(i)
	}
}

func BenchmarkInsRnd(b *testing.B) {
	n := newNode(0, black)
	for i := 0; i < b.N; i++ {
		n, _ = n.insert(rand.Int())
	}
}
