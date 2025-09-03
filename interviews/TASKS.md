## Tasks to split

package main

import "fmt"

func main() {
m := []int{1,2,3,4}
add(m) // 1 2 3 4 5
fmt.Println(m) // 1 2 3 4
add(m) // 1 2 3 4 5
fmt.Println(m) // 1 2 3 4
}

func add(a []int) {
a = append(a, 5)
fmt.Println(a)
}
    
================================================================================    

func main() {
m := make(map[int]int)
m[1] = 10
a := &m[1]
fmt.Println(m[1], *a)
}
    
================================================================================


func foo(m map[int]int) {
m[10] = 10
}

func main() {
m := make(map[int]int)
m[10] = 15
println(m[10]) // 15
foo(m)
println(m[10]) // 10
}

================================================================================

func fn(m map[int]int) {
m = make(map[int]int)
fmt.Println(m == nil) // false
}

func main() {
var m map[int]int
fn(m)
fmt.Println(m == nil) // true
}

================================================================================

func main() {
var wg sync.WaitGroup
wg.Add(10)
for i := 0; i < 10; i++ {
i := i
go func() {
fmt.Println(i)
wg.Done()
}()
}
wg.Wait()
}


================================================================================

Есть 3 сущности: "Автор", "Книга", "Читатель"
Физически книга только 1 и может быть только у одного читателя.
table "author"
id_author INT AUTOINCREMENT
full_name VARCHAR(120)
1
|
m
table "author_book"
id_author
id_book
m
|
1
table "book"
id_book INT AUTOINCREMENT
name VARCHAR(120)
id_reader INT NULL
m
|
1
table "reader"
id_reader INT AUTOINCREMENT
full_name VARCHAR(120)

1) Написать запрос - выбрать названия всех книг которые на руках
   select b.name from book b where b.id_reader is not null;

2)Написать запрос - выбрать названия всех книг в библиотеке у которых больше 3 авторов

select b.name from book b inner join author_book ab on ab.id_book = b.id_book
having count(*) > 3 group by b.id;

================================================================================


Распиши структуру проекта, что куда класть, какие сущность понадобятся.
можно прям там по папочкам расписывать.



cmd/
main.go
internal/
infrastructure/
log
adapters/
postgres/
books.go // books repository
stats.go // stats
app/
books/books.go
stats/stats.go <- StatsGetter impl
ports/
api
handlers/
stats.go -> StatsGetter
v1.go
pkg/
domain/book.go
domain/author.go
