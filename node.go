package cache

type Item interface{}

type Node struct {
	item Item
	next *Node
	prev *Node
}
