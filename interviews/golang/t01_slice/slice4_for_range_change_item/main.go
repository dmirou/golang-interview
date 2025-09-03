// [Annie Betty _]
// Annie
// Betty
// _ (empty string for last item)
// [Annie Jack _]
package main

import "fmt"

func main() {
	friends := [3]string{"Annie", "Betty"}
	fmt.Println(friends) // 1

	for i, v := range friends {
		if i == 0 {
			friends[i+1] = "Jack"
		}
		fmt.Println(v) // 2
	}
	fmt.Println(friends) // 3
}
