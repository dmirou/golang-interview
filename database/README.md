## Links

- http://go-database-sql.org/index.html learn with DB in golang
- SELECT https://postgrespro.ru/docs/postgresql/14/sql-select
- SELECT FOR UPDATE SKIP LOCKED https://www.2ndquadrant.com/en/blog/what-is-select-skip-locked-for-in-postgresql-9-5
- Блокировки https://postgrespro.ru/docs/postgresql/14/explicit-locking, 
- Lock table https://postgrespro.ru/docs/postgresql/9.6/sql-lock 
- Relation level lock https://habr.com/ru/company/postgrespro/blog/500714/ 
- Row level lock https://habr.com/ru/company/postgrespro/blog/503008/ 
- Консистентность данных (foreign key и тд), 
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
- [Postgres 13 documentation](https://www.postgresql.org/docs/13/index.html)
- postgres tutorial https://www.postgresqltutorial.com/


### Transactions

- Уровни изоляции в базе https://postgrespro.ru/docs/postgrespro/9.5/transaction-iso,
  https://habr.com/ru/post/317884/
- Транзакции, точки сохранения https://postgrespro.ru/docs/postgresql/14/tutorial-transactions
- На пути к правильным SQL транзакциям (Часть 1) https://habr.com/ru/company/infopulse/blog/261097/


### Задачи

#### Составление запроса на выборку данных (авторы и книги)

Дана БД:
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

Задача: Выбрать имя авторов, кто опубликовал до 2010 года больше 3х книг

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

#### Индексы 

Добавить индексы, чтобы оптимизировать запрос

```sql
EXPLAIN ANALYSE
SELECT "users"."id",
       "users"."first_name",
       "users"."second_name",
       "users"."last_name",
       "users"."email",
       "users"."address",
       "users"."phone_number",
       "users"."company_id",
       "users"."job_id"
FROM "users"
WHERE (UPPER("users"."id"::text) = UPPER('Josh') OR
       UPPER("users"."first_name"::text) LIKE UPPER('%Josh%') OR
       UPPER("users"."last_name"::text) LIKE UPPER('%Josh%') OR
       UPPER("users"."phone_number"::text) LIKE UPPER('Josh%') OR
       UPPER("users"."email"::text) LIKE UPPER('%Josh%'))
ORDER BY "users"."last_name" ASC;
```
