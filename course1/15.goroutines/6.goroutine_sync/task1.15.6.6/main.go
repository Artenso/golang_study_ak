package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID   int
	Name string
}

// Пример реализации конкурентно-безопасного кэша с использованием sync.Map
type Cache struct {
	data sync.Map
}

func NewCache() *Cache {
	return &Cache{data: sync.Map{}}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.data.Load(key)
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func GetUser(i interface{}) *User {
	return i.(*User)
}

func main() {
	wg := sync.WaitGroup{}

	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprint("user-", i),
		})
	}

	time.Sleep(time.Second)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			raw, ok := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			if !ok {
				fmt.Println("unknown key")
			}
			fmt.Println(GetUser(raw))
		}(i)
	}

	wg.Wait()
}
