package lru

type LRUCache struct {
	size  int
	tail  *Node
	head  *Node
	store map[int]Node
}

type Node struct {
	parent *Node
	child  *Node
	key    int
	value  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size:  capacity,
		store: make(map[int]Node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if record, ok := this.store[key]; !ok {
		return -1
	} else {
		// Update recency
		this.moveToFront(key)
		return record.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	newNode := Node{
		key:   key,
		value: value,
		child: this.head,
	}

	// Add if it doesn't exist
	if _, ok := this.store[key]; !ok {
		this.store[key] = newNode
		this.head.parent = &newNode
		this.head = &newNode
	} else {
		// Update recency
		this.store[key] = newNode
		this.head = &newNode
	}

	// Trim
	if len(this.store) > this.size {
		this.evict()
	}
}

func (this *LRUCache) evict() {
	delete(this.store, this.tail.key)
	this.tail.child = nil
	this.tail = this.tail.parent
}

// Reoder Linked List
func (this *LRUCache) moveToFront(key int) {
	updatedNode := this.store[key]
	// Remove from list
	updatedNode.child.parent = updatedNode.parent
	updatedNode.parent.child = updatedNode.child
	// Attach to head
	updatedNode.child = this.head
	this.head = &updatedNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
