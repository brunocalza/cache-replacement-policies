package cache

type List struct {
	head *Node
	tail *Node
}

func NewList() *List {
	l := &List{}
	l.head = &Node{} //dummy node
	l.tail = &Node{} //dummy node
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *List) AppendLeft(key CacheKey) *Node {
	node := &Node{key, l.head.next, l.head}
	l.head.next.prev = node
	l.head.next = node
	return node
}

func (l *List) Pop() CacheKey {
	temp := l.tail.prev
	l.tail.prev = temp.prev
	temp.prev.next = l.tail
	temp.next = nil
	temp.prev = nil
	return temp.key
}
