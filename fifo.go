package cache

// FIFOPolicy implements the First-In First-Out policy
// The cache evicts the elements in the order that they were added to the cache policer
// It is implemented as a doubly-linked list
type FIFOPolicy struct {
	head    *Node
	tail    *Node
	keyNode map[CacheKey]*Node
}

// NewFIFOPolicy creates a new FIFO cache policer
func NewFIFOPolicy() CachePolicy {
	policy := &FIFOPolicy{}
	policy.head = &Node{} //dummy node
	policy.tail = &Node{} //dummy node
	policy.head.next = policy.tail
	policy.tail.prev = policy.head
	policy.keyNode = make(map[CacheKey]*Node)
	return policy
}

// Victim selects a cache key for eviction using the FIFO policy
func (p *FIFOPolicy) Victim() CacheKey {
	temp := p.tail.prev
	p.tail.prev = temp.prev
	temp.prev.next = p.tail
	temp.next = nil
	temp.prev = nil
	return temp.key
}

// Add adds a cache key to the policer, becoming a candidate for eviction
func (p *FIFOPolicy) Add(key CacheKey) {
	node := &Node{key, p.head.next, p.head}
	p.head.next.prev = node
	p.head.next = node
	p.keyNode[key] = node
}

// Removes a cache key from the policer, so the key is no longer considered for eviction
func (p *FIFOPolicy) Remove(key CacheKey) {
	node, ok := p.keyNode[key]
	if !ok {
		return
	}
	node.next.prev = node.prev
	node.prev.next = node.next
	node.next = nil
	node.prev = nil
}
