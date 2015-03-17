package redblack

const (
	stackSize = 64
)

type elem struct {
	n *node
	d direction
}

type stack struct {
	c uint8
	e [stackSize]elem
}

func (s *stack) push(n *node, d direction) {
	e := &s.e[s.c]
	e.n = n
	e.d = d
	s.c++
}

func (s *stack) pop() (*node, direction) {
	s.c--
	e := &s.e[s.c]
	return e.n, e.d
}
