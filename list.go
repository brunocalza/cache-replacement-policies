package cache

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

func (l *List) AppendLeft(key CacheKey) *Node {
	node := &Node{key, l.head.next, l.head}
	l.head.next.prev = node
	l.head.next = node
	l.size++
	return node
}

func (l *List) Pop() CacheKey {
	temp := l.tail.prev
	l.tail.prev = temp.prev
	temp.prev.next = l.tail
	temp.next = nil
	temp.prev = nil
	l.size--
	return temp.key
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
