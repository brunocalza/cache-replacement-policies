package cache

type Item interface{}

type Node struct {
	item Item
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func NewList() *List {
	l := &List{}
	l.head = &Node{} //dummy node
	l.tail = &Node{} //dummy node
	l.head.next = l.tail
	l.tail.prev = l.head
	l.size = 0
	return l
}

func (l *List) AppendLeft(item Item) *Node {
	node := &Node{item, l.head.next, l.head}
	l.head.next.prev = node
	l.head.next = node
	l.size++
	return node
}

func (l *List) Pop() Item {
	temp := l.tail.prev
	l.tail.prev = temp.prev
	temp.prev.next = l.tail
	temp.next = nil
	temp.prev = nil
	l.size--
	return temp.item
}

func (l *List) Remove(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
	node.next = nil
	node.prev = nil
	l.size--
}

func (l *List) Size() int {
	return l.size
}
