package redblack

type node struct {
	k int
	c color
	x [2]*node
}

func newNode(k int, c color) *node {
	return &node{k: k, c: c, x: [2]*node{nil, nil}}
}

type Tree struct {
	r *node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Root() *node {
	return t.r
}

func (t *Tree) Insert(k int) {
	var s stack
	var n, b, p *node
	var nd, bd, pd direction
	if t.r == nil {
		t.r = newNode(k, black)
		return
	}
	s = stack{}
	n = t.r
	for n != nil {
		switch {
		case n.k > k:
			s.push(n, left)
			n = n.x[left]
		case n.k < k:
			s.push(n, right)
			n = n.x[right]
		default:
			return
		}
	}
	n, nd = s.pop()
	n.x[nd] = newNode(k, red)
	for n.c == red {
		p, pd = s.pop()
		bd = ^pd & 1
		b = p.x[bd]
		if b != nil && b.c == red {
			b.c = black
			n.c = black
			if p != t.r {
				p.c = red
				n, nd = s.pop()
				continue
			}
			return
		}
		if nd != pd {
			n = rotate(n, nd)
			p.x[pd] = n
		}
		p.c = red
		n.c = black
		if p == t.r {
			t.r = rotate(p, pd)
			return
		}
		n, nd = s.pop()
		n.x[nd] = rotate(p, pd)
		return
	}
}

func (t *Tree) Height() int {
	if t.r == nil {
		return 0
	}
	return height(t.r)
}

func rotate(n *node, d direction) *node {
	f := ^d & 1
	a := n.x[d]
	b := a.x[f]
	a.x[f] = n
	n.x[d] = b
	return a
}

func height(n *node) int {
	if n == nil {
		return 1
	}
	a := n.x[left]
	if a != nil {
		if a.c == red && n.c == red {
			return 0
		}
		if a.k >= n.k {
			return 0
		}
	}
	b := n.x[right]
	if b != nil {
		if b.c == red && n.c == red {
			return 0
		}
		if b.k <= n.k {
			return 0
		}
	}
	c := height(a)
	d := height(b)
	if (c != 0) && (d != 0) {
		if c != d {
			return 0
		}
		if n.c == red {
			return c
		}
		return c + 1
	}
	return 0
}
