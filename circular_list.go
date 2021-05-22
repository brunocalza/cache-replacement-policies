package cache

import (
	"container/ring"
	"fmt"
)

type CircularList struct {
	ring *ring.Ring
}

func (c *CircularList) Append(item interface{}) *ring.Ring {
	if c.ring == nil {
		c.ring = ring.New(1)
		c.ring.Value = item
		return c.ring
	}

	newRing := ring.New(1)
	newRing.Value = item
	c.ring.Link(newRing)
	c.ring = newRing
	return newRing
}

func (c *CircularList) Remove(ring *ring.Ring) {
	if c.ring.Len() == 1 {
		if c.ring == ring {
			c.ring = nil
		}
		return
	}

	prev := ring.Prev()
	prev.Unlink(1)
	if ring == c.ring {
		c.ring = prev
	}
}

func (c *CircularList) Move(ring *ring.Ring) {
	c.ring = ring
}

func (c *CircularList) Len() int {
	if c.ring == nil {
		return 0
	}

	return c.ring.Len()
}

func (c *CircularList) Print() {
	if c.ring != nil {
		c.ring.Next().Do(func(p interface{}) {
			fmt.Print(p)
		})
	}
	fmt.Println()
}
