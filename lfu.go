package cache

type Frequency int

type FreqNode struct {
	freq Frequency
	node *Node
}

type LFUPolicy struct {
	keyFreqNode    map[CacheKey]FreqNode
	freqList       map[Frequency]*List
	keyNode        map[CacheKey]*Node
	leastFrequency Frequency
}

// NewLFUPolicy creates a new LFU cache policer
func NewLFUPolicy() CachePolicy {
	policy := &LFUPolicy{}
	policy.keyFreqNode = make(map[CacheKey]FreqNode)
	policy.freqList = make(map[Frequency]*List)
	policy.leastFrequency = 1
	return policy
}

// Victim selects a cache key for eviction using the LFU policy
func (p *LFUPolicy) Victim() CacheKey {
	return p.freqList[p.leastFrequency].Pop()
}

// Add adds a cache key to the policer, becoming a candidate for eviction
func (p *LFUPolicy) Add(key CacheKey) {
	_, ok := p.freqList[1]
	if !ok {
		p.freqList[1] = NewList()
	}

	node := p.freqList[1].AppendLeft(key)
	p.keyFreqNode[key] = FreqNode{1, node}
	p.leastFrequency = 1
}

// Removes a cache key from the policer, so the key is no longer considered for eviction
func (p *LFUPolicy) Remove(key CacheKey) {
	p.remove(key)
}

// Access indicates to the policer that the key was accessed
func (p *LFUPolicy) Access(key CacheKey) {
	freqNode := p.remove(key)

	node := freqNode.node
	frequency := freqNode.freq

	_, ok := p.freqList[frequency+1]
	if !ok {
		p.freqList[frequency+1] = NewList()
	}

	node = p.freqList[frequency+1].AppendLeft(key)
	p.keyFreqNode[key] = FreqNode{frequency + 1, node}
}

func (p *LFUPolicy) remove(key CacheKey) FreqNode {
	freqNode, _ := p.keyFreqNode[key]

	node := freqNode.node
	frequency := freqNode.freq

	p.freqList[frequency].Remove(node)
	delete(p.keyFreqNode, key)

	if p.freqList[frequency].Size() == 0 {
		delete(p.freqList, frequency)
		if p.leastFrequency == frequency {
			p.leastFrequency++
		}
	}

	return freqNode
}
