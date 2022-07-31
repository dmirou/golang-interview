## Commands

### Show linux process memory model

```bash
ps ax | grep main.go
```
12345 - pid процесса
cat /proc/12345/maps | tac


### Get assembler code
```bash
go tool compile -S main.go
```

`// go:noinline` перед функцией, чтобы компилятор это не оптимизировал.

### Escape analysis 

(посмотреть что убегает на кучу)
```bash
go tool compile -m main.go
```
or
```bash
go build -gcflags=-m main.go
```

## Features

- В стеке процесса хранится мета информация для runtime
- Стек горутин хранится в куче процесса, 100 горутин = 100 стеков
- При вызове функции вызывающая функция ложит параметры вызова в стек вызывающей функции, 
и передаёт управление этой функции
- в стеке живут локальные переменные
- глобальные переменные и переменные которые передаются между подпрограммами создаются в куче
- запрос памяти у OS происходит большими фрагментами (arena)
- nmap в go позволяет маппить большой файл в виртуальную память и ходить по ним как будто он в памяти

## Links

- [Understand Go pointers in less than 800 words or your money back](https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back)
- [There is no pass-by-reference in Go](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go)
- https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8
- https://medium.com/a-journey-with-go/go-how-does-the-garbage-collector-mark-the-memory-72cfc12c6976
- Implementing memory management with Golang’s garbage collector 
    https://hub.packtpub.com/implementing-memory-management-with-golang-garbage-collector/
- Garbage Collection In Go : Part I - Semantics
    https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html
- [**Pointers**](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers)
- [Using Pointers In Go](https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html)
- [Language Mechanics On Stacks And Pointers](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html)
- [Go: How Does the Goroutine Stack Size Evolve?](https://medium.com/a-journey-with-go/go-how-does-the-goroutine-stack-size-evolve-447fc02085e5)
- [Go’s hidden #pragmas](https://dave.cheney.net/2018/01/08/gos-hidden-pragmas)