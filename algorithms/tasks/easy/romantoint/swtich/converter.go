package main

import "fmt"

func romanToInt(s string) int {
	var left, right rune
	sum := 0

	length := len(s)

	for i := 0; i < length; {
		left = rune(s[i])

		if i+1 < length {
			right = rune(s[i+1])
		}

		num, both := symbolsToInt(left, right)
		sum += num
		if both {
			i += 2
			continue
		}
		i++
	}

	return sum
}

func symbolsToInt(left, right rune) (num int, both bool) {
	pair := string(left) + string(right)

	both = true
	switch pair {
	case "IV":
		num = 4
		return
	case "IX":
		num = 9
		return
	case "XL":
		num = 40
		return
	case "XC":
		num = 90
		return
	case "CD":
		num = 400
		return
	case "CM":
		num = 900
		return
	default:
		// unknown pair
	}

	both = false
	switch left {
	case 'I':
		num = 1
	case 'V':
		num = 5
	case 'X':
		num = 10
	case 'L':
		num = 50
	case 'C':
		num = 100
	case 'D':
		num = 500
	case 'M':
		num = 1000
	default:
		panic("unknown char " + string(left))
	}

	return
}

func main() {
	var r rune
	fmt.Printf("%v %q\n", r, r)

	r = 44
	fmt.Printf("%c\n", r)
}
