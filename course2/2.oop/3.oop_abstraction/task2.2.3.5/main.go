package main

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"time"

	"github.com/howeyc/crc16"
	"github.com/sigurn/crc8"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type HashMap struct {
	buckets  []bucket
	hashFunc HashFunc
}

type bucket struct {
	pairs map[string]any
}

type HashFunc func(data []byte) uint64

type HashMapOption func(*HashMap)

func NewHashMap(size int, opts ...HashMapOption) HashMaper {
	hm := &HashMap{
		buckets: make([]bucket, size),
	}

	for idx := range hm.buckets {
		hm.buckets[idx] = bucket{pairs: make(map[string]any)}
	}

	for _, opt := range opts {
		opt(hm)
	}

	return hm
}

func WithHashCRC64() HashMapOption {
	return func(hm *HashMap) {
		hm.hashFunc = func(data []byte) uint64 {
			return crc64.Checksum(data, &crc64.Table{})
		}
	}
}

func WithHashCRC32() HashMapOption {
	return func(hm *HashMap) {
		hm.hashFunc = func(data []byte) uint64 {
			return uint64(crc32.Checksum(data, &crc32.Table{}))
		}
	}
}

func WithHashCRC16() HashMapOption {
	return func(hm *HashMap) {
		hm.hashFunc = func(data []byte) uint64 {
			return uint64(crc16.Checksum(data, &crc16.Table{}))
		}
	}
}

func WithHashCRC8() HashMapOption {
	return func(hm *HashMap) {
		hm.hashFunc = func(data []byte) uint64 {
			return uint64(crc8.Checksum(data, &crc8.Table{}))
		}
	}
}

func (hm *HashMap) Set(key string, value interface{}) {
	hash := hm.hashFunc([]byte(key))
	idx := hash % uint64(len(hm.buckets))
	hm.buckets[idx].pairs[key] = value
}

func (hm *HashMap) Get(key string) (interface{}, bool) {
	hash := hm.hashFunc([]byte(key))
	idx := hash % uint64(len(hm.buckets))
	value, ok := hm.buckets[idx].pairs[key]
	return value, ok
}

func MeassureTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}

func main() {
	m := NewHashMap(16, WithHashCRC64())
	since := MeassureTime(func() {
		m.Set("key", "value")
		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since.Nanoseconds())

	m = NewHashMap(16, WithHashCRC32())
	since = MeassureTime(func() {
		m.Set("key", "value")
		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since.Nanoseconds())

	m = NewHashMap(16, WithHashCRC16())
	since = MeassureTime(func() {
		m.Set("key", "value")
		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since.Nanoseconds())

	m = NewHashMap(16, WithHashCRC8())
	since = MeassureTime(func() {
		m.Set("key", "value")
		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since.Nanoseconds())

	m = NewHashMap(16, WithHashCRC64())
	since = MeassureTime(func() {
		m.Set("key", "value")
		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since.Nanoseconds())

}
