package cache

import "container/list"

type Frequency int

type LFUItem struct {
	frequency Frequency
	key       CacheKey
}

type LFUPolicy struct {
	freqList     map[Frequency]*list.List
	keyNode      map[CacheKey]*list.Element
	minFrequency Frequency
}

// NewLFUPolicy creates a new LFU cache policer
func NewLFUPolicy() CachePolicy {
	policy := &LFUPolicy{}
	policy.keyNode = make(map[CacheKey]*list.Element)
	policy.freqList = make(map[Frequency]*list.List)
	policy.minFrequency = 1
	return policy
}

// Victim selects a cache key for eviction using the LFU policy
func (p *LFUPolicy) Victim() CacheKey {
	fList := p.freqList[p.minFrequency]
	element := fList.Back()
	fList.Remove(element)
	delete(p.keyNode, element.Value.(LFUItem).key)
	return element.Value.(LFUItem).key
}

// Add adds a cache key to the policer, becoming a candidate for eviction
func (p *LFUPolicy) Add(key CacheKey) {
	_, ok := p.freqList[1]
	if !ok {
		p.freqList[1] = list.New()
	}

	node := p.freqList[1].PushFront(LFUItem{1, key})
	p.keyNode[key] = node
	p.minFrequency = 1
}

// Removes a cache key from the policer, so the key is no longer considered for eviction
func (p *LFUPolicy) Remove(key CacheKey) {
	p.remove(key)
}

// Access indicates to the policer that the key was accessed
func (p *LFUPolicy) Access(key CacheKey) {
	node := p.remove(key)

	frequency := node.Value.(LFUItem).frequency
	_, ok := p.freqList[frequency+1]
	if !ok {
		p.freqList[frequency+1] = list.New()
	}

	node = p.freqList[frequency+1].PushFront(LFUItem{frequency + 1, key})
	p.keyNode[key] = node
}

func (p *LFUPolicy) remove(key CacheKey) *list.Element {
	node := p.keyNode[key]
	frequency := node.Value.(LFUItem).frequency

	p.freqList[frequency].Remove(node)
	delete(p.keyNode, key)

	if p.freqList[frequency].Len() == 0 {
		delete(p.freqList, frequency)
		if p.minFrequency == frequency {
			p.minFrequency++
		}
	}

	return node
}
