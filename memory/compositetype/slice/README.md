# Slices

## Internals

see src/runtime/slice.go

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

## Declaration

```go
var s1 []int // nil, len is 0, cap is 0
s2 := make([]int, 0, 10) // not nil, len is 0, cap is 10
var s3 = []byte("hello мир")
```

## Coping

To copy slice from `src` you need to create `dst` slice with
the same length or more and call `copy(dst, src)`.
If length of `dst` will be zero or less than `src` only len(dst)
items will be copied. `copy` return number of copied items.

## Features

- Slices hold references to an underlying array, and if you assign one slice to 
    another, both refer to the same array. 
- Slice is passed to the function by value. 
    We must return the slice afterwards because, although Append can modify 
    the elements of slice, the slice itself (the run-time data structure holding 
    the pointer, length, and capacity) is passed by value.
- Slices can be compared only with `nil`. To compare two slices you need to write
    you custom function. It's because underlying array can be changed from another
    go-routine during comparison and also slice can contain himself.
- Go's slices are one-dimensional. To create the equivalent of a 2D slice, 
    it is necessary to define a slice-of-slices:
    ```go
    type LinesOfText [][]byte     // A slice of byte slices.
  
    text := LinesOfText{
    	[]byte("Now is the time"),
    	[]byte("for all good gophers"),
    	[]byte("to bring some fun to the party."),
    }
    ```

## Links

- https://go.dev/doc/effective_go
- https://www.callicoder.com/golang-slices/
- https://www.geeksforgeeks.org/how-to-copy-one-slice-into-another-slice-in-golang/
