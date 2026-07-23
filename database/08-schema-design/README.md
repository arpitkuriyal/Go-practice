# 08. Schema Design: Model Data Before Writing Queries

## Start with the questions

Before creating a table, ask:

1. What real-world thing does one row represent?
2. Which fields are required?
3. Which values must be unique?
4. Which records relate to each other?
5. Which rules must remain true even if application code has a bug?

For example, a post belongs to one user. That relationship should be represented by a foreign key, not by copying the author name into every post.

## Normalization in plain language

Normalization means storing each fact in one sensible place to avoid inconsistent duplicates.

Bad design:

```text
posts: id, title, author_name, author_email
```

If an author changes their email, every old post must be updated too.

Better design:

```text
users: id, name, email
posts: id, author_id, title
```

The post stores the relationship (`author_id`), and a join reads the author details.

## Constraints are part of the design

```sql
CREATE TABLE posts (
    id BIGSERIAL PRIMARY KEY,
    author_id BIGINT NOT NULL REFERENCES users(id),
    title TEXT NOT NULL CHECK (char_length(title) <= 200),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

| Constraint | Business protection |
| --- | --- |
| `PRIMARY KEY` | Every row has a unique identity. |
| `NOT NULL` | Required fields cannot be absent. |
| `UNIQUE` | Values such as email cannot duplicate. |
| `FOREIGN KEY` | A relationship cannot point to a missing record. |
| `CHECK` | A value must satisfy a rule. |

## Choosing an ID

- `BIGSERIAL` is compact and simple for one database.
- `UUID` is useful for public or distributed IDs.
- Never use an email or mutable business value as the primary key.

## When denormalization is okay

Sometimes you intentionally store duplicate, derived data to make a measured read query faster—for example, a cached comment count. This is **denormalization**. Only do it when you know how the duplicate value stays correct.

## Example schema

[`schema-examples.sql`](schema-examples.sql) creates a small normalized blog schema: users create posts, and `user_roles` models a many-to-many user/role relationship. Read the constraints line by line before running it.

## Practical rules

- Store timestamps with time zone as `TIMESTAMPTZ` in PostgreSQL.
- Use integer minor units or a decimal type for money; do not rely on binary floating point.
- Model many-to-many relationships with a join table.
- Use soft deletes only when the product needs recovery/audit behavior; otherwise a normal delete is simpler.
- Design indexes from real query patterns, not only from table shape.

## Interview answer

“I start schema design from entities, required fields, relationships, and invariants. I normalize repeated facts into related tables, enforce critical rules with constraints, and denormalize only for a measured performance reason with a plan to keep duplicate data consistent.”
