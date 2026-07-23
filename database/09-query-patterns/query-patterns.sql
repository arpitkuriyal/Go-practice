-- Filtering: values are placeholders when called from Go.
SELECT id, title, created_at
FROM posts
WHERE author_id = $1
  AND created_at >= $2
ORDER BY created_at DESC, id DESC
LIMIT $3;

-- Offset pagination: simple for small result sets.
SELECT id, title, created_at
FROM posts
ORDER BY created_at DESC, id DESC
LIMIT $1 OFFSET $2;

-- Keyset pagination: use the last row's created_at and id as the cursor.
SELECT id, title, created_at
FROM posts
WHERE (created_at, id) < ($1, $2)
ORDER BY created_at DESC, id DESC
LIMIT $3;

-- Full-text search example for PostgreSQL.
SELECT id, title
FROM posts
WHERE to_tsvector('english', title || ' ' || body)
      @@ plainto_tsquery('english', $1);
