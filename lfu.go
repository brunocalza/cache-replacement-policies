package cache

type Frequency int

type LFUItem struct {
	frequency Frequency
	key       CacheKey
}

type LFUPolicy struct {
	freqList       map[Frequency]*List
	keyNode        map[CacheKey]*Node
	leastFrequency Frequency
}

// NewLFUPolicy creates a new LFU cache policer
func NewLFUPolicy() CachePolicy {
	policy := &LFUPolicy{}
	policy.keyNode = make(map[CacheKey]*Node)
	policy.freqList = make(map[Frequency]*List)
	policy.leastFrequency = 1
	return policy
}

// Victim selects a cache key for eviction using the LFU policy
func (p *LFUPolicy) Victim() CacheKey {
	return p.freqList[p.leastFrequency].Pop().(LFUItem).key
}

// Add adds a cache key to the policer, becoming a candidate for eviction
func (p *LFUPolicy) Add(key CacheKey) {
	_, ok := p.freqList[1]
	if !ok {
		p.freqList[1] = NewList()
	}

	node := p.freqList[1].AppendLeft(LFUItem{1, key})
	p.keyNode[key] = node
	p.leastFrequency = 1
}

// Removes a cache key from the policer, so the key is no longer considered for eviction
func (p *LFUPolicy) Remove(key CacheKey) {
	p.remove(key)
}

// Access indicates to the policer that the key was accessed
func (p *LFUPolicy) Access(key CacheKey) {
	node := p.remove(key)

	frequency := node.item.(LFUItem).frequency
	_, ok := p.freqList[frequency+1]
	if !ok {
		p.freqList[frequency+1] = NewList()
	}

	node = p.freqList[frequency+1].AppendLeft(LFUItem{frequency + 1, key})
	p.keyNode[key] = node
}

func (p *LFUPolicy) remove(key CacheKey) *Node {
	node := p.keyNode[key]
	frequency := node.item.(LFUItem).frequency

	p.freqList[frequency].Remove(node)
	delete(p.keyNode, key)

	if p.freqList[frequency].Size() == 0 {
		delete(p.freqList, frequency)
		if p.leastFrequency == frequency {
			p.leastFrequency++
		}
	}

	return node
}
