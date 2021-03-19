package lru_cache

type LRUCache struct {
	capacity  int
	cache     map[int]*ListItem
	usedOrder *DoubleLinkedList
}


func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		capacity: capacity,
		cache: map[int]*ListItem{},
		usedOrder: &DoubleLinkedList{
			head: &ListItem{},
			tail: &ListItem{},
		},
	}
	lruCache.usedOrder.head.next = lruCache.usedOrder.tail
	lruCache.usedOrder.tail.prev = lruCache.usedOrder.head

	return lruCache
}


func (c *LRUCache) Get(key int) int {
	if el, ok := c.cache[key]; ok {
		c.usedOrder.moveToTail(el)

		return el.value
	}

	return -1
}


func (c *LRUCache) Put(key int, value int)  {
	if el, ok := c.cache[key]; ok {
		el.value = value
		c.usedOrder.moveToTail(el)
	} else {
		// capacity reached, we need to delete LRU key
		// which is in the head
		if len(c.cache) >= c.capacity {
			el := c.usedOrder.head.next
			c.usedOrder.removeElement(el)
			delete(c.cache, el.key)
		}

		// creating new element
		val := &ListItem{key: key, value: value}
		c.usedOrder.addElement(val)
		c.cache[key] = val
	}
}