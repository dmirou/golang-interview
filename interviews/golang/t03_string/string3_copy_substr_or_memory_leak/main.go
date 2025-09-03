// problem:
// return data[i+2 : i+22] doesn't copy the data, it returns pointer to existing substr
// to fix:
// convert substr to slice bytes, it will copy the string because strings are not changeable.
// return string([]byte(data[i+2 : i+22]))
package main

func findNecessaryData(data string) string {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == '\n' && data[i+1] == '\t' {
			return data[i+2 : i+22]
		}
	}

	return ""
}

func main() {
	var data string
	// let's imagine that data was read from file

	necessaryData := findNecessaryData(data)
	_ = necessaryData // use it later
}
