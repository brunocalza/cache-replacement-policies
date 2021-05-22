package cache

import (
	"container/ring"
	"fmt"
)

type CircularList struct {
	ring *ring.Ring
}

func (c *CircularList) Append(item interface{}) *ring.Ring {
	newRing := ring.New(1)
	newRing.Value = item

	if c.ring != nil {
		c.ring.Link(newRing)
	}

	c.ring = newRing // c.ring always points to the last element
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
