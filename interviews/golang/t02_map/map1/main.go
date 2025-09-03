// panic: assignment to entry in nil map
// to fix: add map initialization
// var c = make(Cart)
package main

import "fmt"

type Cart map[string]int32

func (cart Cart) AddOrNew(sku string, count int32) {
	cart[sku] = count
}

func main() {
	var c Cart

	c.AddOrNew("orange", 2)
	c.AddOrNew("orange", 3)

	fmt.Println(&c) //
}
