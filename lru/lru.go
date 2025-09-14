// A LRU cache implementation in go
package lru

// node represents an entry in the doubly linked list.
// It stores a key-value pair and links to its neighbors.
type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

// list is doubly linked list which contains the head,
// tail, size, capacity of the list
type list[K comparable, V any] struct {
	head *node[K, V]
	tail *node[K, V]

	size int
	capacity int
}

func newNode[K comparable, V any] (key K, value V) *node[K, V] {
	return &node[K, V] {
		key: key,
		value: value,
	}
}

func addNode[K comparable, V any](c *LRUCache[K, V], n *node[K, V]) {
	if c.list.capacity == 0 {
		return 
	}

	if c.list.size >= c.list.capacity {
		evict(c)
	}

	if c.list.size == 0 {
		c.list.head = n
		c.list.tail = n
	} else {
		n.next = c.list.head
		c.list.head.prev = n
		c.list.head = n
	}

	c.list.size++ 
}

func evict[K comparable, V any](c *LRUCache[K, V]) {
	if c.list.tail == nil {
		return 
	}

	evicted := c.list.tail

	if evicted.prev != nil {
		c.list.tail = evicted.prev
		c.list.tail.next = nil
		evicted.prev = nil
	} else {
		c.list.head = nil
		c.list.tail = nil
	}

	delete(c.items, evicted.key)
	c.list.size--
}

func moveToFront[K comparable, V any](c *LRUCache[K, V], n *node[K, V]) {
	if n == c.list.head {
		return
	}

	// Detach node fron its current position
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		c.list.tail = n.prev
	}

	// Insert node at head
	n.next = c.list.head
	n.prev = nil
	if c.list.head != nil {
		c.list.head.prev = n
	}
	c.list.head = n

	if c.list.tail == nil {
		c.list.tail = n
	}
}

// LRUCache is a least-recently-used cache.
// It combines a map for fast lookups and a doubly linked list
// to track usage order.
type LRUCache[K comparable, V any] struct {
	items map[K]*node[K, V] // key -> node
	list list[K, V]
}

// Constructor initalizes the LRUCache instance
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V] {
		items: make(map[K]*node[K, V]),
		list: list[K, V] {
			capacity: capacity,
			size: 0,
		},
	}
}


// User API
func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	nodeAddr, exists := c.items[key]
	if !exists {
		var zero V
		return zero, false
	}
	moveToFront(c, nodeAddr)
	return nodeAddr.value, true
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	nodeAddr, exists := c.items[key]
	if exists {
		nodeAddr.value = value
		moveToFront(c, nodeAddr)		
		return
	}

	node := newNode(key, value)
	addNode(c, node)

	c.items[key] = node
}

func (c *LRUCache[K, V]) Capacity() int {
	return c.list.capacity
}

func (c *LRUCache[K, V]) Size() int {
	return c.list.size
}
