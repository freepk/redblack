package redblack

import (
	"testing"
)

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
