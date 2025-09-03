Implement function which will generate numbers from start to finish
with increment inc.

```text
func gen(start, finish, inc int) chan int {
	// TODO implement me
}
func main() {
	for i := range gen(1, 5, 1) {
		fmt.Println(i)
	}
}
```
