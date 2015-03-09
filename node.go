package redblack

const (
	black color = iota
	red
)

type color uint8

type node struct {
	x int
	r color
	l [2]*node
}

func newNode(x int, r color) *node {
	return &node{x: x, r: r, l: [2]*node{nil, nil}}
}

func (n *node) rotate(d direction) *node {
	f := ^d & 1
	x := n.l[d]
	z := x.l[f]
	x.l[f] = n
	n.l[d] = z
	return x
}

func (n *node) insert(x int) (*node, bool) {
	s := newStack()
	if t := s.rewind(n, x); t != nil {
		return n, false
	}
	p, pd := s.pop()
	p.l[pd] = newNode(x, red)
	for p.r == red {
		g, gd := s.pop()
		ud := ^gd & 1
		if u := g.l[ud]; u != nil && u.r == red {
			u.r = black
			p.r = black
			if g != n {
				g.r = red
				p, pd = s.pop()
				continue
			}
		}
		if gd != pd {
			p = p.rotate(pd)
			g.l[gd] = p
		}
		g.r = red
		p.r = black
		switch {
		case g == n:
			return g.rotate(gd), true
		default:
			g = g.rotate(gd)
			t, td := s.pop()
			t.l[td] = g
			return n, true
		}
	}
	return n, true
}
