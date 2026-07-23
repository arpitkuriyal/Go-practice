# 09. Query Patterns: Filtering, Sorting, Search, and Pagination

## Filtering safely

Users often choose filters such as status, category, or date range. Pass values as query arguments:

```go
const query = `SELECT id, title FROM posts WHERE author_id = $1 LIMIT $2`
rows, err := db.QueryContext(ctx, query, authorID, limit)
```

Placeholders protect values. They generally cannot replace SQL identifiers or keywords, so never pass an untrusted column name directly into `ORDER BY`.

## Safe sorting

Choose SQL fragments from a small allow-list:

```go
orderBy := "created_at DESC"
if sort == "title" {
	orderBy = "title ASC"
}
query := "SELECT id, title FROM posts ORDER BY " + orderBy + " LIMIT $1"
```

Only the application-controlled allow-list is concatenated; user values remain parameters.

## Pagination

Every list endpoint should have a limit. Choose a sensible default and maximum.

### Offset pagination

```sql
SELECT id, title
FROM posts
ORDER BY created_at DESC, id DESC
LIMIT $1 OFFSET $2;
```

It is simple and good for small lists, but deep offsets become slower and inserts can shift later pages.

### Cursor (keyset) pagination

```sql
SELECT id, title, created_at
FROM posts
WHERE (created_at, id) < ($1, $2)
ORDER BY created_at DESC, id DESC
LIMIT $3;
```

The cursor is the last item's `created_at` and `id`. This is efficient for deep, changing lists when supported by a matching index.

## Search

- For an exact value, use `WHERE email = $1` and a normal index.
- For a prefix, `LIKE 'arpit%'` may be appropriate with the right index.
- For document/word search, use PostgreSQL full-text search.
- Do not use `LIKE '%text%'` on a large table without understanding its cost and available index options.

## Dynamic query checklist

- Validate every filter type and range.
- Parameterize every value.
- Allow-list dynamic SQL pieces such as sort fields and direction.
- Apply a stable order before pagination.
- Cap the page size.
- Select only columns the endpoint needs.

## Interview answer

“I parameterize filter values, allow-list any dynamic sort clause, and always use a stable order plus a bounded limit. Offset pagination is simple; keyset pagination is better for deep or frequently changing result sets.”

## Examples

[`query-patterns.sql`](query-patterns.sql) contains copyable PostgreSQL examples for filtering, offset pagination, keyset pagination, and full-text search. `$1`, `$2`, and `$3` are placeholders used by a PostgreSQL driver from Go; replace them with values when experimenting directly in a SQL client.
