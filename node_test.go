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
	if isValid(n) == 0 {
		t.Fail()
	}
}

func insertSeq(c int) *node {
	n := newNode(0, black)
	for i := 0; i < c; i++ {
		n, _ = n.insert(i)
	}
	return n
}

func insertRnd(c int) *node {
	n := newNode(0, black)
	for i := 0; i < c; i++ {
		n, _ = n.insert(rand.Int())
	}
	return n
}

func BenchmarkInsSeq(b *testing.B) {
	insertSeq(b.N)
}

func BenchmarkInsRnd(b *testing.B) {
	insertRnd(b.N)
}

func BenchmarkSearchSeq(b *testing.B) {
	b.N = 1000000
	n := insertSeq(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.search(i)
	}
}
