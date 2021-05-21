package cache

import "container/list"

type LRUPolicy struct {
	list    *list.List
	keyNode map[CacheKey]*list.Element
}

// NewLRUPolicy creates a new LRU cache policer
func NewLRUPolicy() CachePolicy {
	policy := &LRUPolicy{}
	policy.list = list.New()
	policy.keyNode = make(map[CacheKey]*list.Element)
	return policy
}

// Victim selects a cache key for eviction using the LRU policy
func (p *LRUPolicy) Victim() CacheKey {
	element := p.list.Back()
	p.list.Remove(element)
	return element.Value.(CacheKey)
}

// Add adds a cache key to the policer, becoming a candidate for eviction
func (p *LRUPolicy) Add(key CacheKey) {
	node := p.list.PushFront(key)
	p.keyNode[key] = node
}

// Removes a cache key from the policer, so the key is no longer considered for eviction
func (p *LRUPolicy) Remove(key CacheKey) {
	node, ok := p.keyNode[key]
	if !ok {
		return
	}
	p.list.Remove(node)
	delete(p.keyNode, key)
}

// Access indicates to the policer that the key was accessed
func (p *LRUPolicy) Access(key CacheKey) {
	p.Remove(key)
	p.Add(key)
}
