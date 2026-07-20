# Query Performance: Revision

## Diagnose before changing

Use the real query and real-like data:

```sql
EXPLAIN ANALYZE
SELECT id, title
FROM posts
WHERE author_id = $1
ORDER BY created_at DESC
LIMIT 20;
```

Read the plan for sequential scans, row estimates, actual time, sort operations, and expensive joins. Do not add indexes only because a column “looks important.”

## Index rules

- Index columns commonly used in `WHERE`, joins, and selective ordering.
- A composite index follows leftmost-prefix rules. `(author_id, created_at DESC)` helps filtering by `author_id` and ordering by `created_at`.
- Every index consumes storage and slows inserts/updates/deletes.
- An index is usually unnecessary for a low-cardinality boolean column by itself.

```sql
CREATE INDEX posts_author_created_at_idx
    ON posts (author_id, created_at DESC);
```

## Common interview traps

- **N+1 query:** fetch a list, then issue one query per row. Fix with a join, batch query, or preloading strategy.
- **Offset pagination:** `OFFSET` gets slower on deep pages and can shift under writes. Use keyset/cursor pagination for large, changing datasets.
- **`SELECT *`:** transfers and couples more data than needed. Select only required columns.
- **Missing limits:** a list endpoint should have an explicit, bounded page size.

## Keyset pagination example

```sql
SELECT id, title, created_at
FROM posts
WHERE (created_at, id) < ($1, $2)
ORDER BY created_at DESC, id DESC
LIMIT $3;
```
