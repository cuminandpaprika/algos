package lru

type LRUCache struct {
	size      int
	filled    int
	nextEvict *Node
	newest    *Node
	store     map[int]int
}

type Node struct {
	child *Node
	key   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size:  capacity,
		store: make(map[int]int, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.store[key]; !ok {
		return -1
	} else {
		// Update recency
		this.push(key)
		return val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.store[key]; !ok {
		this.store[key] = value
	} else {
		// Update recency
		this.push(key)
		this.filled++
		this.store[key] = value
	}
	if this.filled > this.size {
		this.evict()
	}
}

func (this *LRUCache) evict() {
	if _, ok := this.store[this.nextEvict.key]; !ok {
		this.nextEvict = this.nextEvict.child
		this.evict()
	} else {
		delete(this.store, this.nextEvict.key)
		this.nextEvict = this.nextEvict.child
	}

}

func (this *LRUCache) push(key int) {
	latest := &Node{
		key: key,
	}

	if this.newest != nil {
		this.newest.child = latest
	}

	this.newest = latest
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
