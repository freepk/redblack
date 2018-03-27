package redblack

type color byte

type direction byte

const (
	red color = iota
	black
)

const (
	left direction = iota
	right
)

const (
	stackSize = 64
)
