package main

import "sync"

type SafeMap struct {
	Data map[string]interface{}
	Lock sync.RWMutex
}

func (sm *SafeMap) SetData(key string, data interface{}) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()
	if sm.Data == nil {
		sm.Data = make(map[string]interface{})
	}
	sm.Data[key] = data
}

func (sm *SafeMap) GetData(key string) interface{} {
	sm.Lock.RLock()
	defer sm.Lock.RUnlock()

	if value, exists := sm.Data[key]; exists {
		return value
	}
	return nil
}
