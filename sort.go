package main

import (
	"fmt"
	"sort"
	// "strconv"
)

type Fruit struct {
	name  string
	color string
	price float64
}

func (f Fruit) String() string {
	return fmt.Sprintf("%s ($%v)", f.name, f.price)
}

type FruitBasket []Fruit

func (b FruitBasket) Len() int              { return len(b) }
func (b FruitBasket) Swap(i, j int)         { b[i], b[j] = b[j], b[i] }
func (b FruitBasket) Less(i, j int) bool    { return b[i].name < b[j].name }
func (b FruitBasket) More(i, j int) bool    { return b[i].name > b[j].name }
func (b FruitBasket) Get(i int) interface{} { return b[i] }


type ByPrice []Fruit

func (b ByPrice) Len() int           { return len(b) }
func (b ByPrice) Swap(i, j int)		 { b[i], b[j] = b[j], b[i] }
func (b ByPrice) Less(i, j int) bool { return b[i].price < b[j].price }
func (b ByPrice) More(i, j int) bool { return b[i].price > b[j].price }
func (b ByPrice) Get(i int) interface{} { return b[i] }

// MaxInterface is an interface defining the methods that must be 
// implemented by any argument passed to the Max function.
type MaxInterface interface {
    Len() int               // number of elements
    Get(i int) interface{}	// elem with index i
    More(i, j int) bool     // elem at index i is > j?
}

// Max returns the maximum value in the collection.
func Max(data MaxInterface) interface{} {
    if data.Len() == 0 {
        return nil
    }
    if data.Len() == 1 {
        return data.Get(1)
    }
    max := data.Get(0)
    m := 0
    for i := 1; i < data.Len(); i++ {
        if data.More(i, m) {
            max = data.Get(i)
            m = i
        }
    }
    return max
}

// MinInterface is an interface defining the methods that must be 
// implemented by any argument passed to the Min function.
type MinInterface interface {
    Len() int               // number of elements
    Get(i int) interface{}	// elem with index i
    Less(i, j int) bool     // elem at index i is < j?
}

// Min returns the minimum value in the collection.
func Min(data MinInterface) interface{} {
    if data.Len() == 0 {
        return nil
    }
    if data.Len() == 1 {
        return data.Get(1)
    }
    min := data.Get(0)
    m := 0
    for i := 1; i < data.Len(); i++ {
        if data.Less(i, m) {
            min = data.Get(i)
            m = i
        }
    }
    return min
}

func main() {
	b := FruitBasket{
		Fruit{"banana", "yellow", .75},
		Fruit{"apple", "red", 1.05},
		Fruit{"orange", "orange", 0.99},
	}

	fmt.Println(b)
	sort.Sort(b)
	fmt.Println(b)
	sort.Sort(ByPrice(b))
	fmt.Println(b)
	fmt.Println(Max(ByPrice(b)))
	fmt.Println(Min(ByPrice(b)))
}
