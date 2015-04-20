package redblack

import (
	"math/rand"
	"testing"
)

var (
	sample *node
)

func init() {
	sample = fillSeq(10000000)
}

func fillSeq(n int) *node {
	x := newNode(0, black)
	for i := 0; i < n; i++ {
		x, _ = x.insert(i)
	}
	return x
}

func fillRnd(n int) *node {
	x := newNode(0, black)
	for i := 0; i < n; i++ {
		x, _ = x.insert(rand.Int())
	}
	return x
}

func TestCommon(t *testing.T) {
	x := fillRnd(10000)
	if isValid(x) == 0 {
		t.Fail()
	}
}

func BenchmarkInsSeq(b *testing.B) {
	fillSeq(b.N)
}

func BenchmarkInsRnd(b *testing.B) {
	fillRnd(b.N)
}

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sample.search(rand.Int())
	}
}
