// What's wrong?
// 1. It's better to read file by chunks
// 2. To compare byte and rune you need to convert byte to rune, e.g. rune(data[i]) == 'M'
// 3. You should copy result into the new memory to allow GC clean big buffer
// result := make([]byte, 20)
// copy(result, data[i:i+20])
// return result
package main

func FindData(filename string) []byte {
	data := make([]byte, 1<<30)
	// read big file into data

	for i := 0; i < len(data)-1; i++ {
		if data[i] == 'М' && data[i+1] == 'А' {
			result := make([]byte, 20)
			copy(result, data[i:i+20])
			return data[i : i+20]
		}
	}
}
