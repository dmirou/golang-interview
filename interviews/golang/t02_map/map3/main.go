package main

import "fmt"

type Cart map[string]int32

func (cart Cart) AddOrNew(sku string, count int32) {
	cart[sku] = count
}

func (cart Cart) Get(sku string) int32 {
	// TODO implement me
	return 0
}

func (cart Cart) Get2(sku string) (int32, error) {
	// TODO implement me
	return 0, nil
}

func main() {
	var c = make(Cart)
	c.AddOrNew("orange", 2)

	fmt.Println(c.Get("orange"))  // expected: 2
	fmt.Println(c.Get("blue"))    // expected: 0
	fmt.Println(c.Get2("orange")) // expected: 2 nil
	fmt.Println(c.Get2("blue"))   // expected: 0 not found
}
