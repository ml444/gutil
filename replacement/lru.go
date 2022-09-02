//go:build go1.9
// +build go1.9

package replacement

import (
	"container/list"
	"sync"
)

type LRU struct {
	size      int
	innerList *list.List
	// innerMap  map[int]*list.Element
	innerMap sync.Map
}

type entry struct {
	key   int
	value int
}

func (lru *LRU) Get(key int) (int, bool) {
	if e, ok := lru.innerMap.Load(key); ok {
		el, ok := e.(*list.Element)
		if !ok {
			return -1, false
		}
		lru.innerList.MoveToFront(el)
		return el.Value.(*entry).value, true
	}
	return -1, false
}

func (lru *LRU) Put(key int, value int) (evicted bool) {
	if e, ok := lru.innerMap.Load(key); ok {
		el := e.(*list.Element)
		lru.innerList.MoveToFront(el)
		el.Value.(*entry).value = value
		return false
	} else {
		e := &entry{key, value}
		el := lru.innerList.PushFront(e)
		lru.innerMap.Store(key, el)

		if lru.innerList.Len() > lru.size {
			last := lru.innerList.Back()
			lru.innerList.Remove(last)
			lru.innerMap.Delete(last.Value.(*entry).key)
			return true
		}
		return false
	}
}




