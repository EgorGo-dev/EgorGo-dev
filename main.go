package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func TotalPrice(product []Product) float64 {
	var total float64
	for _, p := range product {
		total += p.Price
	}
	return total
}

func MostExpensive(product []Product) (Product, bool) {
	if len(product) == 0 {
		return Product{}, false
	}

	best := product[0]
	for _, p := range product {
		if p.Price > best.Price {
			best = p
		}
	}

	return best, true
}

func main() {
	products := []Product{
    {Name: "Книга", Price: 500},
    {Name: "Ручка", Price: 50},
    {Name: "Ноутбук", Price: 50000},
}

	fmt.Println(TotalPrice(products))          // 50550
	p, ok := MostExpensive(products)
	fmt.Println(p.Name, p.Price, ok)  
}