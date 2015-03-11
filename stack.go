package redblack

const (
	stackSize = 64
)

const (
	left direction = iota
	right
)

type direction uint8

type elem struct {
	n *node
	d direction
}

type stack struct {
	c uint8
	e [stackSize]elem
}

func newStack() stack {
	return stack{c: 0}
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

func (s *stack) rewind(n *node, x int) *node {
	s.c = 0
	for n != nil {
		switch {
		case n.x > x:
			s.push(n, left)
			n = n.l[left]
		case n.x < x:
			s.push(n, right)
			n = n.l[right]
		default:
			return n
		}
	}
	return nil
}

func (s *stack) items() []elem {
	return s.e[:s.c]
}
