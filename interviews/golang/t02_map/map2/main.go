// map[]
// map[orange:6]
// map[orange:6]
package main

import "fmt"

type Cart map[string]int32

func (cart Cart) AddOrNew(sku string, count int32) {
	cart[sku] = count
}

func main() {
	var cart = make(Cart, 1)
	fmt.Println(cart) // ?
	var d = &cart

	cart.AddOrNew("orange", 2)
	(*d).AddOrNew("orange", 1)
	(*d)["orange"] += 5

	fmt.Println(cart) // ?
	fmt.Println(*d)   // ?
}
