package utils

import (
	"fmt"
	"miner/models"
	"sync"
)

type Key string

type Value map[string]interface{}

type ValueHashtable struct {
	Items map[string]Value
	lock  sync.RWMutex
}

func hash(k Key) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

func (ht *ValueHashtable) Put(k string, v models.StatResult) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	if ht.Items == nil {
		ht.Items = make(map[string]Value)
	}
	if ht.Items[k] == nil {
		ht.Items[k] = make(map[string]interface{})
	}
	hashElem := ht.Items[k]
	hashElem[v.Name] = v
	ht.Items[k] = hashElem
}

func (ht *ValueHashtable) Remove(k string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	delete(ht.Items, k)
}

func (ht *ValueHashtable) Get(k string) Value {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return ht.Items[k]
}

func (ht *ValueHashtable) Size() int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return len(ht.Items)
}

var resultHash ValueHashtable

func GetResultHash() *ValueHashtable {
	return &resultHash
}

func Init() {
	resultHash = ValueHashtable{}
}
