package cache

import (
	"container/ring"
)

type ClockItem struct {
	key CacheKey
	bit bool
}

type ClockPolicy struct {
	list      *CircularList
	keyNode   map[CacheKey]*ring.Ring
	clockHand *ring.Ring
}

// NewClockPolicy creates a new Clock cache policer
func NewClockPolicy() CachePolicy {
	policy := &ClockPolicy{}
	policy.keyNode = make(map[CacheKey]*ring.Ring)
	policy.list = &CircularList{}
	policy.clockHand = nil
	return policy
}

// Victim selects a cache key for eviction using the Clock policy
func (p *ClockPolicy) Victim() CacheKey {
	var victimKey CacheKey
	var nodeItem *ClockItem
	for {
		currentNode := (*p.clockHand)
		nodeItem = currentNode.Value.(*ClockItem)
		if nodeItem.bit {
			nodeItem.bit = false
			currentNode.Value = nodeItem
			p.clockHand = currentNode.Next()
		} else {
			victimKey = nodeItem.key
			p.list.Move(p.clockHand.Prev())
			p.clockHand = nil
			p.list.Remove(&currentNode)
			delete(p.keyNode, victimKey)
			return victimKey
		}
	}
}

// Add adds a cache key to the policer, becoming a candidate for eviction
func (p *ClockPolicy) Add(key CacheKey) {
	node := p.list.Append(&ClockItem{key, true})
	if p.clockHand == nil { // it means that either the list was empty or a element was just removed
		p.clockHand = node
	}
	p.keyNode[key] = node
}

// Removes a cache key from the policer, so the key is no longer considered for eviction
func (p *ClockPolicy) Remove(key CacheKey) {
	node, ok := p.keyNode[key]
	if !ok {
		return
	}

	if p.clockHand == node {
		p.clockHand = p.clockHand.Prev()
	}
	p.list.Remove(node)
	delete(p.keyNode, key)
}

// Access indicates to the policer that the key was accessed
func (p *ClockPolicy) Access(key CacheKey) {
	node, ok := p.keyNode[key]
	if !ok {
		return
	}
	node.Value = &ClockItem{key, true}
}
