package cache

import (
	"errors"
)

type CacheKey string

type CacheData map[CacheKey]string

type CachePolicy interface {
	Victim() CacheKey
	Add(CacheKey)
	Remove(CacheKey)
}

type Cache struct {
	maxSize int
	size    int
	policy  CachePolicy
	data    CacheData
}

func NewCache(maxSize int, policy CachePolicy) *Cache {
	cache := &Cache{}
	cache.maxSize = maxSize
	cache.policy = policy
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
		return &value, nil
	}

	return nil, errors.New("key not found")
}
