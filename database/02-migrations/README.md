# 02. Migrations: Change a Database Safely

## The beginner problem

At first, you create a table manually. Later, your application needs a new column or table. Every developer, test environment, and production server must make the same change in the same order.

A **migration** is a version-controlled SQL file that records one schema change.

```text
001_create_users.up.sql
002_add_user_avatar.up.sql
003_create_posts.up.sql
```

Think of migrations as Git commits for database structure.

## What is a schema?

The **schema** is the shape of the database: tables, columns, constraints, indexes, and relationships.

This folder’s [`001_create_users_and_posts.up.sql`](migrations/001_create_users_and_posts.up.sql) creates two tables:

```text
users  1 ──── many  posts
```

Each post stores `author_id`, which points to one user.

## Up and down migrations

| File | Purpose |
| --- | --- |
| `001_create_users_and_posts.up.sql` | Apply the change: create tables and an index. |
| `001_create_users_and_posts.down.sql` | Undo the change: remove the tables. |

The `up` migration moves forward. A `down` migration can help during development, but production rollbacks need careful planning because deleting a column or table can delete data.

## Constraints in the example

```sql
id BIGSERIAL PRIMARY KEY
email TEXT NOT NULL UNIQUE
author_id BIGINT NOT NULL REFERENCES users(id)
```

| Constraint | Protects against |
| --- | --- |
| Primary key | Rows without a stable unique ID. |
| `NOT NULL` | Missing required data. |
| `UNIQUE` | Duplicate email addresses. |
| Foreign key | A post pointing to a user that does not exist. |

`ON DELETE CASCADE` in this example means deleting a user also deletes that user’s posts. Choose this only when it matches the product rule.

## Safe migration rules

1. Give migrations ordered, descriptive names.
2. Once a migration has reached shared or production environments, never edit it—add a new migration instead.
3. Run migrations through CI/CD or a dedicated migration job, not from every app instance at startup.
4. Back up and test destructive changes.
5. For a large live table, use an expand/contract rollout: add compatible structure, deploy code, backfill data, then remove old structure later.

## Next step

When one business operation requires several SQL statements to succeed together, use a transaction → [03 Transactions](../03-transactions-and-locking/README.md).

## Interview answer

“Migrations are ordered, version-controlled schema changes. They keep every environment on the same database version. I never edit an already deployed migration; I add a new one and plan destructive changes carefully.”
