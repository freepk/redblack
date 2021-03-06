package redblack

import (
	"math/rand"
	"testing"
)

func benchmarkTreeSeq(n int) *Tree {
	t := NewTree()
	for i := 0; i < n; i++ {
		t.Insert(i)
	}
	return t
}

func benchmarkTreeRnd(n int) *Tree {
	t := NewTree()
	for i := 0; i < n; i++ {
		t.Insert(rand.Int())
	}
	return t
}

func TestNewStack(t *testing.T) {
	s := stack{}
	s.push(newNode(10, red), left)
	s.push(newNode(20, black), right)
	n, d := s.pop()
	if n.k != 20 || n.c != black || d != right {
		t.Fail()
	}
	n, d = s.pop()
	if n.k != 10 || n.c != red || d != left {
		t.Fail()
	}
}

func BenchmarkStackPushAndPop(b *testing.B) {
	s := stack{}
	for i := 0; i < b.N; i++ {
		switch (i / stackSize) % 2 {
		case 0:
			s.push(nil, left)
		case 1:
			s.pop()
		}
	}
}

func TestTreeCommon(t *testing.T) {
	var tr *Tree
	tr = benchmarkTreeSeq(10000)
	if height(tr.r) == 0 {
		t.Fail()
	}
	tr = benchmarkTreeRnd(10000)
	if height(tr.r) == 0 {
		t.Fail()
	}
}

func BenchmarkTreeInsSeq(b *testing.B) {
	benchmarkTreeSeq(b.N)
}

func BenchmarkTreeInsRnd(b *testing.B) {
	benchmarkTreeRnd(b.N)
}

func sampleNode() *node {
	/*

	             40(b)
	       20(r)       60(r)
	   10(b) 30(b) 50(b) 70(b)

	*/

	n10 := newNode(10, black)
	n30 := newNode(30, black)
	n50 := newNode(50, black)
	n70 := newNode(70, black)
	n20 := newNode(20, red)
	n60 := newNode(60, red)
	n40 := newNode(40, black)

	n40.x[left] = n20
	n40.x[right] = n60
	n20.x[left] = n10
	n20.x[right] = n30
	n60.x[left] = n50
	n60.x[right] = n70

	return n40
}

func TestNodeRotate(t *testing.T) {
	n := sampleNode()

	if height(n) != 3 {
		t.Fail()
	}
	n = rotate(n, right)
	if n.k != 60 {
		t.Fail()
	}
	n = rotate(n, left)
	if n.k != 40 {
		t.Fail()
	}
	if height(n) != 3 {
		t.Fail()
	}
	n = rotate(n, left)
	if n.k != 20 {
		t.Fail()
	}
	n = rotate(n, right)
	if n.k != 40 {
		t.Fail()
	}
	if height(n) != 3 {
		t.Fail()
	}
}

func BenchmarkNodeRotate(b *testing.B) {
	n := sampleNode()
	for i := 0; i < b.N; i++ {
		n = rotate(n, right)
		n = rotate(n, left)
	}
}
