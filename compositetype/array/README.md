# Arrays

- Arrays are values. Assigning one array to another copies all the elements.
- In particular, if you pass an array to a function, it will receive a copy 
    of the array, not a pointer to it.
- The size of an array is part of its type. The types [10]int and [20]int are distinct.
- You can compare two arrays between each other via == and !=
- Go's arrays are one-dimensional. To create the equivalent of a 2D array, 
    it is necessary to define an array-of-arrays:
    ```go
    type Transform [3][3]float64  // A 3x3 array, really an array of arrays.
    ```

## Links

- https://go.dev/doc/effective_go