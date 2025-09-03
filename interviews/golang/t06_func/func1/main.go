// 1 foo
// 2 bar
// 1 baz
// ERROR: undefined: j and s2
package main

func main() {
	var i, s1 = 1, "foo"

	{
		var j, s2 = 2, "bar"

		println(i, s1) // 1
		println(j, s2) // 2

		s1 = "baz"
	}

	println(i, s1) // 3
	println(j, s2) // 4
}
