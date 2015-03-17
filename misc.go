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
