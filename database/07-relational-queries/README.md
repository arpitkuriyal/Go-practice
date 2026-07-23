# 07. Relational Queries: Joins, Aggregates, and `NULL`

This lesson builds on the `users` and `posts` relationship from the migrations lesson.

## Why joins exist

Relational databases keep related data in separate tables instead of copying it repeatedly.

```text
users                         posts
+----+-------+                 +----+-----------+--------------+
| id | name  |                 | id | author_id | title        |
+----+-------+                 +----+-----------+--------------+
| 1  | Arpit |                 | 10 | 1         | Learn Go     |
+----+-------+                 +----+-----------+--------------+
```

`posts.author_id` points to `users.id`. A `JOIN` lets one query combine those related rows.

## `INNER JOIN`: only matching rows

```sql
SELECT posts.id, posts.title, users.name AS author_name
FROM posts
INNER JOIN users ON users.id = posts.author_id;
```

Read it as: “for every post, find its user where the IDs match.” An inner join returns only rows with a match on both sides.

## `LEFT JOIN`: keep every row from the left table

```sql
SELECT users.name, posts.title
FROM users
LEFT JOIN posts ON posts.author_id = users.id;
```

This returns every user, including a user with no posts. For that user, `posts.title` is `NULL`.

| Join | Use it when |
| --- | --- |
| `INNER JOIN` | You only want related records that exist on both sides. |
| `LEFT JOIN` | You want all rows from the left table, even without a match. |

## Aggregates: summarize rows

| Function | Meaning |
| --- | --- |
| `COUNT(*)` | Count rows. |
| `SUM(amount)` | Add numeric values. |
| `AVG(salary)` | Calculate an average. |
| `MIN` / `MAX` | Find the smallest / largest value. |

`GROUP BY` makes one result per group:

```sql
SELECT users.id, users.name, COUNT(posts.id) AS post_count
FROM users
LEFT JOIN posts ON posts.author_id = users.id
GROUP BY users.id, users.name
ORDER BY post_count DESC;
```

`WHERE` filters rows **before** grouping. `HAVING` filters groups **after** grouping:

```sql
SELECT author_id, COUNT(*) AS post_count
FROM posts
GROUP BY author_id
HAVING COUNT(*) >= 3;
```

## `NULL`: unknown or absent, not an empty value

`NULL` does not mean `0`, `""`, or `false`. It means the value is absent or unknown.

```sql
-- This is wrong: equality with NULL is unknown, not true.
SELECT * FROM users WHERE deleted_at = NULL;

-- This is correct.
SELECT * FROM users WHERE deleted_at IS NULL;
```

Use `COALESCE` to choose a fallback in a result:

```sql
SELECT name, COALESCE(bio, 'No bio yet') AS bio
FROM users;
```

In Go, scan nullable database columns into types such as `sql.NullString`, or model optional values carefully with pointers.

## Many-to-many relationships

If users can have many roles and a role can belong to many users, use a join table:

```text
users ← user_roles → roles
```

Do not store a comma-separated list of role IDs in one column; it is hard to constrain and query.

## Practice

See [relational-queries.sql](relational-queries.sql) for queries to run after the migrations.

## Interview answer

“A join combines rows through related keys. I use an inner join when both records must exist and a left join when I must keep unmatched left-side rows. `GROUP BY` aggregates rows, and I handle `NULL` explicitly with `IS NULL`, `COALESCE`, and nullable Go scan types.”
