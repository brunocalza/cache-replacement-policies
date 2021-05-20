package cache

// FIFOPolicy implements the First-In First-Out policy
// The cache evicts the elements in the order that they were added to the cache policer
// It is implemented as a doubly-linked list
type FIFOPolicy struct {
	list    *List
	keyNode map[CacheKey]*Node
}

// NewFIFOPolicy creates a new FIFO cache policer
func NewFIFOPolicy() CachePolicy {
	policy := &FIFOPolicy{}
	policy.list = NewList()
	policy.keyNode = make(map[CacheKey]*Node)
	return policy
}

// Victim selects a cache key for eviction using the FIFO policy
func (p *FIFOPolicy) Victim() CacheKey {
	return p.list.Pop().(CacheKey)
}

// Add adds a cache key to the policer, becoming a candidate for eviction
func (p *FIFOPolicy) Add(key CacheKey) {
	node := p.list.AppendLeft(key)
	p.keyNode[key] = node
}

// Removes a cache key from the policer, so the key is no longer considered for eviction
func (p *FIFOPolicy) Remove(key CacheKey) {
	node, ok := p.keyNode[key]
	if !ok {
		return
	}
	p.list.Remove(node)
	delete(p.keyNode, key)
}

// Access indicates to the policer that the key was accessed
// Since accessing the cache has no effect on the eviction policy it just returns
func (p *FIFOPolicy) Access(key CacheKey) {
}
