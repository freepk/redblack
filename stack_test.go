package redblack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := new(stack)
	s.push(newNode(10, 1), left)
	s.push(newNode(20, 1), left)
	n0, _ := s.pop()
	n1, _ := s.pop()
	if n0.x != 20 || n1.x != 10 {
		t.Fail()
	}
}

func BenchmarkStackPushAndPop(b *testing.B) {
	s := new(stack)
	for i := 0; i < b.N; i++ {
		switch (i / stackSize) % 2 {
		case 0:
			s.push(nil, left)
		case 1:
			s.pop()
		}
	}
}
