package main

type RedisBuilder interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}
