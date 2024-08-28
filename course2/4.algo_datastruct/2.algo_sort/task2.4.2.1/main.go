package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/brianvoe/gofakeit"
)

type Product struct {
	Name      string
	Price     float64
	CreatedAt time.Time
	Count     int
}

func (p Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %f, Count: %v", p.Name, p.Price, p.Count)
}

func generateProducts(n int) []Product {
	gofakeit.Seed(time.Now().UnixNano())

	products := make([]Product, n)
	for i := range products {
		products[i] = Product{
			Name:      gofakeit.Word(),
			Price:     gofakeit.Price(1.0, 100.0),
			CreatedAt: gofakeit.Date(),
			Count:     gofakeit.Number(1, 100),
		}
	}

	return products
}

type ProductsSlice struct {
	L        int
	LessFunc func(i, j int) bool
	SwapFunc func(i, j int)
}

func (p *ProductsSlice) Len() int {
	return p.L
}

func (p *ProductsSlice) Less(i, j int) bool {
	return p.LessFunc(i, j)
}

func (p *ProductsSlice) Swap(i, j int) {
	p.SwapFunc(i, j)
}

func ByPrice(products []Product) sort.Interface {
	return &ProductsSlice{
		L:        len(products),
		LessFunc: func(i, j int) bool { return products[i].Price < products[j].Price },
		SwapFunc: func(i, j int) { products[i], products[j] = products[j], products[i] },
	}
}

func ByCreatedAt(products []Product) sort.Interface {
	return &ProductsSlice{
		L:        len(products),
		LessFunc: func(i, j int) bool { return products[i].CreatedAt.Before(products[j].CreatedAt) },
		SwapFunc: func(i, j int) { products[i], products[j] = products[j], products[i] },
	}
}

func ByCount(products []Product) sort.Interface {
	return &ProductsSlice{
		L:        len(products),
		LessFunc: func(i, j int) bool { return products[i].Count < products[j].Count },
		SwapFunc: func(i, j int) { products[i], products[j] = products[j], products[i] },
	}
}

func main() {
	products := generateProducts(10)
	fmt.Println("Исходный список:")
	fmt.Println(products)

	// Сортировка продуктов по цене
	sort.Sort(ByPrice(products))
	fmt.Println("\nОтсортировано по цене:")
	fmt.Println(products)

	// Сортировка продуктов по дате создания
	sort.Sort(ByCreatedAt(products))
	fmt.Println("\nОтсортировано по дате создания:")
	fmt.Println(products)

	// Сортировка продуктов по количеству
	sort.Sort(ByCount(products))
	fmt.Println("\nОтсортировано по количеству:")
	fmt.Println(products)
}
