package cache

import "container/list"

// FIFOPolicy implements the First-In First-Out policy
// The cache evicts the elements in the order that they were added to the cache policer
// It is implemented as a doubly-linked list
type FIFOPolicy struct {
	list    *list.List
	keyNode map[CacheKey]*list.Element
}

// NewFIFOPolicy creates a new FIFO cache policer
func NewFIFOPolicy() CachePolicy {
	policy := &FIFOPolicy{}
	policy.list = list.New()
	policy.keyNode = make(map[CacheKey]*list.Element)
	return policy
}

// Victim selects a cache key for eviction using the FIFO policy
// removes the last element of the queue
func (p *FIFOPolicy) Victim() CacheKey {
	element := p.list.Back()
	p.list.Remove(element)
	delete(p.keyNode, element.Value.(CacheKey))
	return element.Value.(CacheKey)
}

// Add adds a cache key to the policer, becoming a candidate for eviction
func (p *FIFOPolicy) Add(key CacheKey) {
	node := p.list.PushFront(key)
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
