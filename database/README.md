## Authors and books

/*
authors                             
+---------------+---------+         
| id            | int     |<-------+
| name          | varchar |        |
+---------------+---------+        |
                                   |
books                              |   author_books
+---------------+---------+        |  +---------------+---------+
| id            | int     |<----+  +--| author_id     | int     |
| name          | varchar |     +-----| book_id       | int     |
| public_year   | int2    |           +---------------+---------+
+---------------+---------+
*/
-- Выбрать имя авторов, кто опубликовал до 2010 года больше 3х книг

```sql
SELECT a.name FROM authors a WHERE a.id IN (
    SELECT author_id FROM author_books ab 
        INNER JOIN books b on ab.book_id = b.id and b.public_year < 2010 
    GROUP BY author_id 
    HAVING COUNT(*) > 3
    )
```


```sql
select a.name, count(ab.book_id) as count_book
from authors as a
inner join author_books as ab on a.id = ab.author_id 
inner join books as b on ab.book_id = b.id 
where b.public_year < 2010
group by a.id
having count(ab.book_id) > 3
```
