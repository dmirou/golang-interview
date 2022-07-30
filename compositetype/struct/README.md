# Struct

## Receivers

- The rule about pointers vs. values for receivers is that value methods can be 
invoked on pointers and values, but pointer methods can only be invoked on pointers.
- Invoking a pointer method on a value inserts the address operator automatically.

## Links

- https://go.dev/doc/effective_go#methods
- https://go.dev/doc/effective_go#embedding
- [Structure size optimization in Golang (alignment/padding). More effective memory layout 
(linters)](https://itnext.io/structure-size-optimization-in-golang-alignment-padding-more-effective-memory-layout-linters-fffdcba27c61)

