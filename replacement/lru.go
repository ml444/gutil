package replacement

import "container/list"

type LRU struct {
	size      int
	innerList *list.List
	innerMap  map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func (lru *LRU) Get(key int) (int, bool) {
	if e, ok := lru.innerMap[key]; ok {
		lru.innerList.MoveToFront(e)
		return e.Value.(*entry).value, true
	}
	return -1, false
}

func (lru *LRU) Put(key int, value int) (evicted bool) {
	if e, ok := lru.innerMap[key]; ok {
		lru.innerList.MoveToFront(e)
		e.Value.(*entry).value = value
		return false
	} else {
		e := &entry{key, value}
		el := lru.innerList.PushFront(e)
		lru.innerMap[key] = el

		if lru.innerList.Len() > lru.size {
			last := lru.innerList.Back()
			lru.innerList.Remove(last)
			delete(lru.innerMap, last.Value.(*entry).key)
			return true
		}
		return false
	}
}
