package lru_cache

type DoubleLinkedList struct {
	head *ListItem
	tail *ListItem
}

type ListItem struct {
	key   int
	value int
	next  *ListItem
	prev  *ListItem
}

func (l *DoubleLinkedList) addElement(el *ListItem) {
	el.prev = l.tail.prev
	el.next = l.tail

	l.tail.prev.next = el
	l.tail.prev = el
}

func (l *DoubleLinkedList) removeElement(el *ListItem) {
	prev := el.prev
	next := el.next

	prev.next = next
	next.prev = prev
}

func (l *DoubleLinkedList) moveToTail(el *ListItem) {
	l.removeElement(el)
	l.addElement(el)
}
