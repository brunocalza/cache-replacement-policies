package cache

type LRUPolicy struct {
	list    *List
	keyNode map[CacheKey]*Node
}

// NewLRUPolicy creates a new LRU cache policer
func NewLRUPolicy() CachePolicy {
	policy := &LRUPolicy{}
	policy.list = NewList()
	policy.keyNode = make(map[CacheKey]*Node)
	return policy
}

// Victim selects a cache key for eviction using the LRU policy
func (p *LRUPolicy) Victim() CacheKey {
	return p.list.Pop()
}

// Add adds a cache key to the policer, becoming a candidate for eviction
func (p *LRUPolicy) Add(key CacheKey) {
	node := p.list.AppendLeft(key)
	p.keyNode[key] = node
}

// Removes a cache key from the policer, so the key is no longer considered for eviction
func (p *LRUPolicy) Remove(key CacheKey) {
	node, ok := p.keyNode[key]
	if !ok {
		return
	}
	node.next.prev = node.prev
	node.prev.next = node.next
	node.next = nil
	node.prev = nil
}

// Access indicates to the policer that the key was accessed
func (p *LRUPolicy) Access(key CacheKey) {
	p.Remove(key)
	p.Add(key)
}
