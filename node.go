package cache

type Node struct {
	key  CacheKey
	next *Node
	prev *Node
}
