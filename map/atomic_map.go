package main

import "sync/atomic"

type AtomicMap struct {
	Data   map[interface{}]interface{}
	IsSafe int64
}

const IsUse int64 = 0
const UnUse int64 = 1

func (c AtomicMap) SetZero(val int64) {
	atomic.StoreInt64(&c.IsSafe, val)
}

func (c AtomicMap) GetAtomic() int64 {
	return atomic.LoadInt64(&c.IsSafe)
}

func (c AtomicMap) SetMap(key interface{}, val interface{}) {

}

func (c AtomicMap) Get(key interface{}) interface{} {
	return nil
}
