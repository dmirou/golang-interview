## Authors and books

```
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
```

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

## Links

- SELECT https://postgrespro.ru/docs/postgresql/14/sql-select
- SELECT FOR UPDATE SKIP LOCKED https://www.2ndquadrant.com/en/blog/what-is-select-skip-locked-for-in-postgresql-9-5
- Блокировки https://postgrespro.ru/docs/postgresql/14/explicit-locking, 
- Lock table https://postgrespro.ru/docs/postgresql/9.6/sql-lock 
- Relation level lock https://habr.com/ru/company/postgrespro/blog/500714/ 
- Row level lock https://habr.com/ru/company/postgrespro/blog/503008/ 
- консистентность данных (foreign key и тд), 
- уровни изоляции в базе https://postgrespro.ru/docs/postgrespro/9.5/transaction-iso, 
https://habr.com/ru/post/317884/ 
- Транзакции, точки сохранения https://postgrespro.ru/docs/postgresql/14/tutorial-transactions 
- Модификаторы целостности данных (constraints) https://postgrespro.ru/docs/postgresql/14/ddl-constraints 
- B дерево - структура, как используются в базах данных, Формирование дерева
- Бд - master/slave, зачем нужно, откуда читать данные, можно ли читать из мастера, если да, то зачем
- Чем отличается асинхронная и синхронная реплика https://www.delphiplus.org/obespechenie-vysokoi-dostupnosti-sistem-na-osnove/smysl-asinkhronnoi-replikatsii.html 
- Основные индексы в базе (BTREE, HASH, Gin, Gist, ROM)
- Хеш таблицы, как работают таблицы - индексы, 
- полнотекстовый индекс https://postgrespro.ru/docs/postgresql/14/textsearch 
- Применимость индексов к разным наборам данных, 
- Партицирование и шардирование, разница
- Cross join, декартово умножение 2 таблиц
- Нормализация данных https://info-comp.ru/database-normalization 
- http://go-database-sql.org/index.html learn with DB in golang
