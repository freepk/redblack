package redblack

import (
	"fmt"
)

func traverse(n *node) {
	if n.l[left] != nil {
		traverse(n.l[left])
	}
	fmt.Printf("ptr: %p node: %v\n", n, n)
	if n.l[right] != nil {
		traverse(n.l[right])
	}
}

func isRed(n *node) bool {
	return (n != nil) && (n.c == red)
}

func isValid(n *node) int {
	if n == nil {
		return 1
	} else {
		ln := n.l[left]
		rn := n.l[right]
		if isRed(n) {
			if isRed(ln) || isRed(rn) {
				return 0
			}
		}
		lh := isValid(ln)
		rh := isValid(rn)
		if (ln != nil && ln.x >= n.x) || (rn != nil && rn.x <= n.x) {
			return 0
		}
		if (lh != 0) && (rh != 0) && (lh != rh) {
			return 0
		}
		if (lh != 0) && (rh != 0) {
			if isRed(n) {
				return lh
			} else {
				return lh + 1
			}
		} else {
			return 0
		}
	}
}
