package hw04lrucache

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

type Key string

var _ Cache = (*lruCache)(nil)

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear() error                        // Очистить кэш
	printCash() []interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*listItem
	mux      *sync.Mutex
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if l.capacity == 0 {
		return false
	}
	l.mux.Lock()
	defer l.mux.Unlock()
	if _, ok := l.items[key]; ok {
		l.items[key].Value.(*Item).value = value
		if err := l.queue.MoveToFront(l.items[key]); err == nil {
			l.items[key] = l.queue.Front()
			return ok
		}
	} else {
		l.queue.PushFront(&Item{
			key:   key,
			value: value,
		})
		l.items[key] = l.queue.Front()
		if l.queue.Len() > l.capacity {
			if err := l.Clear(); err != nil {
				fmt.Printf("%v\r\n", errors.Wrap(err, "can't Set element "+string(key)))
			}
		}
		return ok
	}
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if _, ok := l.items[key]; ok {
		if err := l.queue.MoveToFront(l.items[key]); err == nil {
			l.items[key] = l.queue.Front()
			return l.items[key].Value.(*Item).value, ok
		}
	} else {
		return nil, ok
	}
	return nil, false
}

func (l *lruCache) Clear() error {
	lastItem := l.queue.Back()
	if lastItem == nil {
		return errors.New("queue is empty")
	}
	if item, ok := lastItem.Value.(*Item); ok {
		delete(l.items, item.key)
		if err := l.queue.Remove(lastItem); err != nil {
			return errors.Wrap(err, "can't clear lruCache")
		}
	}
	return nil
}

type Item struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	if capacity < 0 {
		panic("capacity value must be >= 0")
	}
	cash := &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    map[Key]*listItem{},
		mux:      &sync.Mutex{},
	}
	return cash
}

func (l *lruCache) printCash() []interface{} {
	l.mux.Lock()
	defer l.mux.Unlock()
	sl := []interface{}{}
	elem := l.queue.Back()
	for i := 0; i < l.queue.Len(); i++ {
		sl = append(sl, elem.Value.(*Item).key)
		elem = elem.Next
	}
	return sl
}
