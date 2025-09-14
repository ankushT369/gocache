// A LRU cache implementation in go
package lru

// Node represents an entry in the doubly linked list.
// It stores a key-value pair and links to its neighbors.
type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

// List is doubly linked list which contains the head,
// tail, size, capacity of the list
type List[K comparable, V any] struct {
	head *Node[K, V]
	tail *Node[K, V]

	size int
	capacity int

	// to be used for later implementation
	freePtr *Node[K, V]// optional pointer for node reuse
}

func NewNode[K comparable, V any] (key K, value V) *Node[K, V] {
	return &Node[K, V] {
		key: key,
		value: value,
	}
}

func FillList(c *LRUCache, node *node) {
	// enviction policy should be implemented later
	if c.list.size >= c.list.capacity {
		fmt.Println("ENVICTION not implemented");
		return 
	}

	if c.list.size == 0 {
		c.list.head = node;
		c.list.tail = node;
	} else {
		c.list.tail.next = node;	
		node.prev = c.list.tail;
		c.list.tail = node;
	}

	c.list.size++;
}

// LRUCache is a least-recently-used cache.
// It combines a map for fast lookups and a doubly linked list
// to track usage order.
type LRUCache[K comparable, V any] struct {
	items map[K]*node[K, V] // key -> node
	list List[K, V]
}

// Constructor initalizes the LRUCache instance
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V] {
		items: make(map[K]*node[K, V]),
		list: List[K, V] {
			capacity: capacity,
			size: 0,
		},
	}
}

func (c *LRUCache) Capacity() int {
	return c.list.capacity;
}

func (c *LRUCache) Size() int {
	return c.list.size;
}

func Get() {

}

func Put[K comparable, V any]() {

}
