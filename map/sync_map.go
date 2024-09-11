package main

import "sync"

var GoSafeMap sync.Map

func SetGoSafeMap(key, value interface{}) {
	GoSafeMap.Store(key, value)
}

func GetGoSafeMap(key string) interface{} {
	if data, exists := GoSafeMap.Load(key); exists {
		return data
	}
	return nil
}
