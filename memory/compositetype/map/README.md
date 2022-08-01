# Maps

## Internal

see runtime/map.go

map is a pointer to hmap structure.

```go
// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}

// A bucket for a Go map.
type bmap struct {
	// tophash generally contains the top byte of the hash value
	// for each key in this bucket. If tophash[0] < minTopHash,
	// tophash[0] is a bucket evacuation state instead.
	tophash [bucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt elems.
	// NOTE: packing all the keys together and then all the elems together makes the
	// code a bit more complicated than alternating key/elem/key/elem/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}
```

## Usage

```go
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
}
```

### Get map entry

An attempt to fetch a map value with a key that is not present in the map will 
return the zero value for the type of the entries in the map.
```go
timezone["MSK] // 0 (zero value for int)
seconds, ok = timeZone["MSK"] // seconds will be 0, ok = false
seconds, ok = timeZone["UTC"] // seconds will be 0*60*60, ok = true
```

### Delete map entry

To delete a map entry, use the delete built-in function, whose arguments are the 
map and the key to be deleted. It's safe to do this even if the key is already absent 
from the map.

```go
delete(timeZone, "PDT")  // Now on Standard Time
```

## Features

- The key can be of any type for which the equality operator is defined, 
    such as integers, floating point and complex numbers, strings, pointers, 
    interfaces (as long as the dynamic type supports equality), structs and arrays. 
    Slices cannot be used as map keys, because equality is not defined on them. 

- Like slices, maps hold references to an underlying data structure. 
    If you pass a map to a function that changes the contents of the map, 
    the changes will be visible in the caller.
    ```go
    func foo(m map[int]int) { 
        m[10] = 10 
    }
    
    func main() {
    	m := make(map[int]int)
    	m[10] = 15
    	println("m[10] before foo =", m[10]) // 15
    	foo(m)
    	println("m[10] after foo =", m[10]) // 10
    }
    ```
    Значение поменялось. «Что же, мапа передается по ссылке?», — спросите вы. **Нет.** 
    **В Go не бывает ссылок. Невозможно создать 2 переменные с 1 адресом,** 
    как в С++ например. Но зато **можно создать 2 переменные, указывающие на один адрес** 
    (**это указатели**).
    
    ```go
    func fn(m map[int]int) {
    	m = make(map[int]int)
    	fmt.Println("m == nil in fn?:", m == nil) // false
    }
    
    func main() {
    	var m map[int]int
    	fn(m)
    	fmt.Println("m == nil in main?:", m == nil) // true
    }
    ```
     Переменная m передалась **по значению**, поэтому, как в случае с передачей в функцию 
     обычного int, не поменялась (поменялась локальная копия значения в fn). 
     Тогда почему же меняется значение, лежащее в самой m? 
     Потому что **мапа в Go — это просто указатель на структуру hmap**. 
     Это и является ответом на вопрос, почему при том, что мапа передается в функцию по 
     значению, сами значения, лежащие в ней меняются — все дело в указателе. 

## Links

- https://go.dev/doc/effective_go
- https://habr.com/ru/post/457728/
- https://www.youtube.com/watch?v=Tl7mi9QmLns&ab_channel=GopherAcademy
