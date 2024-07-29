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

// Структура Cache содержит хранилище и мьютекс
// для хранения объектов любого типа используем пустой интерфейс
type Cache struct {
	mu   sync.RWMutex
	data map[interface{}]interface{}
}

// Создаем новый объект Cache
func NewCache() *Cache {
	return &Cache{data: make(map[interface{}]interface{})}
}

// Для безопасной записи в хранилище используем блокировку
func (c *Cache) Set(key string, user *User) {
	c.mu.Lock()
	c.data[key] = user
	c.mu.Unlock()
}

// Для безопасного чтения используем болкировку для чтения
func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if user, ok := c.data[key]; ok {
		return user
	}
	return nil
}

// Объединяем ключи с помощью Join из пакета strings
func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

// Для получения указателя на структуру User после получения данных из хранилища
// используем приведение типов, преобразуя пустой интерфейс к типу
// указатель на User
func GetUser(i interface{}) *User {
	return i.(*User)
}

func main() {
	wg := sync.WaitGroup{}

	cache := NewCache() // Создаем Cache

	// Генерируем структуры User и добавляем в кеш
	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprint("user-", i),
		})
	}

	time.Sleep(time.Second) // ждем пока все пользователи будут занесены в кеш

	// Получаем данные из кеша и приводим к типу указатель на пользователя,
	// после чего печатаем полученный результат
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			raw := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			fmt.Println(GetUser(raw))
		}(i)
	}

	// Дожидаемся завершения всех горутин (чтобы все результаты были напечатаны)
	wg.Wait()
}
