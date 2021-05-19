package cache

import (
	"errors"
)

type CacheKey string

type CacheData map[CacheKey]string

type Cache struct {
	maxSize int
	size    int
	policy  CachePolicy
	data    CacheData
}

type PolicyType int

const (
	FIFO PolicyType = 1 << iota
	LRU
	LFU
	CLOCK
)

type CachePolicy interface {
	Victim() CacheKey
	Add(CacheKey)
	Remove(CacheKey)
	Access(CacheKey)
}

func GetCachePolicy(policy PolicyType) CachePolicy {
	switch policy {
	case FIFO:
		return NewFIFOPolicy()
	case LRU:
		return NewLRUPolicy()
	case LFU:
		return NewLFUPolicy()
	default:
		return NewFIFOPolicy()
	}
}

func NewCache(maxSize int, policy PolicyType) *Cache {
	cache := &Cache{}
	cache.maxSize = maxSize
	cache.policy = GetCachePolicy(policy)
	cache.data = make(CacheData, maxSize)
	return cache
}

func (c *Cache) Put(key CacheKey, value string) error {
	if c.size == c.maxSize {
		victimKey := c.policy.Victim()
		delete(c.data, victimKey)
		c.size--
	}
	c.policy.Add(key)
	c.data[key] = value
	c.size++
	return nil
}

func (c *Cache) Get(key CacheKey) (*string, error) {
	if value, ok := c.data[key]; ok {
		c.policy.Access(key)
		return &value, nil
	}

	return nil, errors.New("key not found")
}
