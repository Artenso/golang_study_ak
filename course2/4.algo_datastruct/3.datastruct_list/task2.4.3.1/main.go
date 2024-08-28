package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brianvoe/gofakeit"
)

var _ LinkedLister = &DoubleLinkedList{}

type LinkedLister interface {
	LoadData(path string) error
	Init(c []Commit)
	Len() int
	SetCurrent(n int) error
	Current() *Node
	Next() *Node
	Prev() *Node
	Insert(n int, c Commit) error
	Push(c Commit) error
	Delete(n int) error
	DeleteCurrent() error
	Index() (int, error)
	GetByIndex(n int) (*Node, error)
	Pop() *Node
	Shift() *Node
	SearchUUID(uuID string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

type Commit struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
	Date    string `json:"date"`
}

type Node struct {
	data *Commit
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node // начальный элемент в списке
	tail *Node // последний элемент в списке
	curr *Node // текущий элемент меняется при использовании методов next, prev
	len  int   // количество элементов в списке
}

func (d *DoubleLinkedList) Init(commits []Commit) {
	for _, commit := range commits {
		if err := d.Push(commit); err != nil {
			log.Fatalf("failed to init DoubleLinkedList: %s", err)
		}
	}
}

// SetCurrent не очень понял, но судя по всему нужно установить в качестве текущего элемент,
// у которго порядковый номер от начала списка равен n...
func (d *DoubleLinkedList) SetCurrent(n int) error {
	if n > d.len || n <= 0 {
		return fmt.Errorf("index out of range")
	}

	// Посмотрим откуда ближе до n от начала или от конца
	if n <= d.len/2 {
		d.curr = d.head

		for i := 1; i != n; i++ {
			d.Next()
		}
	} else {

		d.curr = d.tail

		for i := d.len; i != n; i-- {
			d.Prev()
		}
	}

	return nil
}

func (d *DoubleLinkedList) Push(c Commit) error {
	newNode := &Node{
		data: &c,
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		d.len++
	} else {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}

	return nil
}

// GetByIndex получает элемент из списка по порядковому номеру от начала списка
func (d *DoubleLinkedList) GetByIndex(n int) (*Node, error) {
	if n > d.len || n <= 0 {
		return nil, fmt.Errorf("index out of range")
	}

	res := &Node{}

	// Посмотрим откуда ближе до n от начала или от конца
	if n <= d.len/2 {
		d.curr = d.head

		for i := 1; i != n; i++ {
			res = d.Next()
		}

	} else {
		d.curr = d.tail

		for i := d.len; i != n; i-- {
			res = d.Prev()
		}
	}

	return res, nil
}

// LoadData loads data from a JSON file at the given path into the list.
func (d *DoubleLinkedList) LoadData(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var commits []Commit
	if err := json.Unmarshal(bytes, &commits); err != nil {
		return err
	}
	QuickSort(commits)
	d.Init(commits)
	return nil
}

// Len получение длины списка
func (d *DoubleLinkedList) Len() int {
	return d.len
}

// Current получение текущего элемента
func (d *DoubleLinkedList) Current() *Node {
	return d.curr
}

// Next получение следующего элемента
func (d *DoubleLinkedList) Next() *Node {
	if d.curr != nil {
		d.curr = d.curr.next
	}

	return d.curr
}

// Prev получение предыдущего элемента
func (d *DoubleLinkedList) Prev() *Node {
	if d.curr != nil {
		d.curr = d.curr.prev
	}

	return d.curr
}

// Insert писал не я

// Insert вставка элемента после n элемента

// Insert inserts a new node with commit c at position n.
func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n < 0 || n > d.len {
		return errors.New("index out of bounds")
	}
	newNode := &Node{data: &c}
	if n == 0 {
		if d.head == nil {
			d.head = newNode
			d.tail = newNode
		} else {
			newNode.next = d.head
			d.head.prev = newNode
			d.head = newNode
		}
	} else if n == d.len {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	} else {
		current := d.head
		for i := 0; i < n; i++ {
			current = current.next
		}
		newNode.next = current
		newNode.prev = current.prev
		current.prev.next = newNode
		current.prev = newNode
	}
	d.len++
	return nil
}

// Delete удаление n элемента
func (d *DoubleLinkedList) Delete(n int) error {
	if n > d.len || n <= 0 {
		return fmt.Errorf("index out of range")
	}

	// если надо удалить первый элемнт
	if n == 1 {
		d.head = d.head.next
		d.head.prev = nil

		d.curr = d.head

		return nil
	}

	// если надо удалить последний
	if n == d.len {
		d.tail = d.tail.prev
		d.tail.next = nil

		d.curr = d.tail

		return nil
	}

	// если надо удалить любой другой
	deleted := &Node{}

	// посмотрим откуда ближе до n от начала или от конца
	if n <= d.len/2 {
		d.curr = d.head

		for i := 1; i != n; i++ {
			deleted = d.Next()
		}
	} else {

		d.curr = d.tail

		for i := d.len; i != n; i-- {
			deleted = d.Prev()
		}
	}

	// удалим все ссылки на удаляемый объект
	deleted.prev.next = deleted.next
	deleted.next.prev = deleted.prev

	// сделаем текущим элемент следующий после удаленного
	d.curr = deleted.next

	return nil
}

// DeleteCurrent удаление текущего элемента
func (d *DoubleLinkedList) DeleteCurrent() error {
	// если curr сейчас nil
	if d.curr == nil {
		return fmt.Errorf("current is nil")
	}

	// если curr - первый элемент
	if d.curr == d.head {
		d.head = d.head.next
		d.head.prev = nil

		d.curr = d.head

		return nil
	}

	// если curr - последний элемент
	if d.curr == d.tail {
		d.tail = d.tail.prev
		d.tail.next = nil

		d.curr = d.tail

		return nil
	}

	// если curr - любой другой
	d.curr.prev.next = d.curr.next
	d.curr.next.prev = d.curr.prev

	d.curr = d.curr.next

	return nil
}

// Index получение индекса текущего элемента
func (d *DoubleLinkedList) Index() (int, error) {
	// если curr сейчас nil
	if d.curr == nil {
		return 0, fmt.Errorf("current is nil")
	}

	current := d.curr
	d.curr = d.head

	index := 1

	for d.curr != current {
		d.Next()
		index++
	}

	return index, nil
}

// Pop Операция Pop
func (d *DoubleLinkedList) Pop() *Node {
	if d.tail == nil {
		return nil
	}

	res := d.tail

	d.tail = d.tail.prev
	d.tail.next = nil

	return res
}

// Shift операция shift
func (d *DoubleLinkedList) Shift() *Node {
	if d.head == nil {
		return nil
	}

	res := d.head

	d.head = d.head.next
	d.head.prev = nil

	return res
}

// SearchUUID поиск коммита по uuid
func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	if d.head == nil {
		return nil
	}

	d.curr = d.head

	for i := 1; i <= d.len; i++ {
		if d.curr.data.UUID == uuID {
			return d.curr
		}
		d.Next()
	}

	return nil
}

// Search поиск коммита по message
func (d *DoubleLinkedList) Search(message string) *Node {
	if d.head == nil {
		return nil
	}

	d.curr = d.head

	for i := 1; i <= d.len; i++ {
		if d.curr.data.Message == message {
			return d.curr
		}
		d.Next()
	}

	return nil
}

// Reverse возвращает перевернутый список
func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	reversedList := &DoubleLinkedList{}

	for i := d.len; i > 0; i-- {
		node := d.Pop()

		if reversedList.head == nil {
			reversedList.head = node
			reversedList.tail = node
			reversedList.len++
		} else {
			node.next = d.head
			reversedList.head.prev = node
			reversedList.head = node
		}
	}

	return reversedList
}

func QuickSort(commits []Commit) {
	if len(commits) < 2 {
		return
	}

	left, right := 0, len(commits)-1

	pivotIndex := len(commits) / 2

	commits[pivotIndex], commits[right] = commits[right], commits[pivotIndex]

	for i := range commits {
		date1, _ := time.Parse("2006-01-02", commits[i].Date)
		date2, _ := time.Parse("2006-01-02", commits[right].Date)
		if date1.Before(date2) {
			commits[i], commits[left] = commits[left], commits[i]
			left++
		}
	}

	commits[left], commits[right] = commits[right], commits[left]

	QuickSort(commits[:left])
	QuickSort(commits[left+1:])
}

func GenerateData(numCommits int) []Commit {
	var commits []Commit
	gofakeit.Seed(0) // Initialize the random seed
	// Define how many commits you want to generate
	for i := 0; i < numCommits; i++ {
		commit := Commit{
			Message: gofakeit.Sentence(5), // Generate a random sentence with 5 words
			UUID:    gofakeit.UUID(),      // Generate a random UUID
			Date: gofakeit.DateRange( // Generate a random date between 2020 and 2022
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
			).Format("2006-01-02"),
		}
		commits = append(commits, commit)
	}
	return commits
}
