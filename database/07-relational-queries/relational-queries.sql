-- Assumes the users and posts tables from 02-migrations exist.

-- Every post with its author's name.
SELECT posts.id, posts.title, users.name AS author_name
FROM posts
INNER JOIN users ON users.id = posts.author_id;

-- Every user, including users who have not written a post.
SELECT users.id, users.name, COUNT(posts.id) AS post_count
FROM users
LEFT JOIN posts ON posts.author_id = users.id
GROUP BY users.id, users.name
ORDER BY post_count DESC, users.id;

-- Only authors with at least three posts.
SELECT author_id, COUNT(*) AS post_count
FROM posts
GROUP BY author_id
HAVING COUNT(*) >= 3;
