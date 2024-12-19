package converter

import (
	"fmt"
	"testing"
)

type Book struct {
	Name  string
	Price float64
}
type Item struct {
	Name  string
	Price float64
}
type Cart struct {
	TotalPrice float64
	Qty        int
	Item       interface{}
}

func TestConvert(t *testing.T) {
	book := Book{Name: "The Kubernetes Bible", Price: 56.99}
	item := Item{}

	err := Convert(book, &item)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Name:  %s | Price: %f\n", item.Name, item.Price)

}
