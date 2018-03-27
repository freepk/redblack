package redblack

type stack struct {
	c int
	n [stackSize]*node
	d [stackSize]direction
}

func (s *stack) push(n *node, d direction) {
	c := s.c
	s.c++
	s.n[c] = n
	s.d[c] = d
}

func (s *stack) pop() (*node, direction) {
	s.c--
	return s.n[s.c], s.d[s.c]
}
