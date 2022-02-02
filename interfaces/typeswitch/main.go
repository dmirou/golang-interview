package main

import "fmt"

func main() {
	var m = map[string]interface{}{
		"i":        10,
		"str":      "hello",
		"bool":     true,
		"float":    1.15,
		"[]string": []string{""},
	}

	for k, v := range m {
		switch v := v.(type) {
		case int, float32:
			fmt.Printf("%s is int of float(type is %T): %d\n", k, v, v)
		case bool:
			fmt.Printf("%s is bool(type is %T): %t\n", k, v, v)
		case string:
			fmt.Printf("%s is string(type is %T): %s\n", k, v, v)
		default:
			fmt.Printf("%s is unknown(type is %T): %v\n", k, v, v)
		}
	}
}
