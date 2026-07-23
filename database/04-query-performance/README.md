# 04. Query Performance: Make SQL Faster for the Right Reason

 : what makes a query slow?

When a table has only ten rows, almost any query feels fast. As it grows, the database may need to inspect many rows, sort them, or join them with another table.

Do not guess. First measure the real query with realistic data:

```sql
EXPLAIN ANALYZE
SELECT id, title
FROM posts
WHERE author_id = 42
ORDER BY created_at DESC
LIMIT 20;
```

`EXPLAIN ANALYZE` asks PostgreSQL to run the query and show its plan and timing.

## What is an index?

An index is like the index at the back of a book. Instead of checking every page (row), the database can quickly find likely matching rows.

```sql
CREATE INDEX posts_author_created_at_idx
    ON posts (author_id, created_at DESC);
```

This index helps the query above because it filters by `author_id` and then reads rows in `created_at DESC` order.

## Index rules for beginners

- Index columns frequently used in `WHERE`, joins, and important ordering.
- An index makes writes a little more expensive because it must also be updated.
- Do not index every column.
- A composite index `(author_id, created_at)` is ordered from left to right; it helps queries starting with `author_id`.
- Check the query plan and workload before adding an index.

## Common API problems

### `SELECT *`

```sql
-- Better: return only what the client needs.
SELECT id, title, created_at FROM posts;
```

### Missing `LIMIT`

Never let a list endpoint accidentally return an unlimited number of rows. Parse and cap the page size.

### N+1 queries

```text
1 query to list 100 posts
100 more queries to fetch each author
```

Fix it with a join, batch query, or preloading approach.

### Deep offset pagination

`OFFSET 100000` makes the database skip many rows. For large changing data, use cursor/keyset pagination:

```sql
SELECT id, title, created_at
FROM posts
WHERE (created_at, id) < ($1, $2)
ORDER BY created_at DESC, id DESC
LIMIT $3;
```

## Interview answer

“I diagnose slow queries with `EXPLAIN ANALYZE` before changing them. I add indexes that match real filter and ordering patterns, select only needed columns, bound list results, and avoid N+1 queries.”
