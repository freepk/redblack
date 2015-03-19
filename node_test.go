package redblack

import (
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	n := insertSeq(1000)
	if h := isValid(n); h == 0 {
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
	n := insertSeq(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.search(i)
	}
}
