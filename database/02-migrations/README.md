# Migrations: Revision

Migrations are ordered, version-controlled schema changes. Each deployed application version must know exactly which schema version it expects.

## Rules

- Give each migration an ordered, descriptive name: `001_create_users.up.sql`.
- Keep an `up` migration small and focused; provide a safe `down` migration when practical.
- Run migrations through CI/CD or a dedicated migration job—not every application instance at startup.
- Never edit a migration that has been deployed. Add a new migration to correct it.
- Back up and test destructive changes. For large tables, use an expand/contract rollout: add new schema, deploy compatible code, backfill, then remove old schema later.

## Constraints to know

- `PRIMARY KEY` uniquely identifies a row.
- `FOREIGN KEY` maintains relationships.
- `NOT NULL`, `UNIQUE`, and `CHECK` enforce data correctness close to the data.
- Application validation improves user feedback; database constraints protect integrity under every writer.

The sample migration creates `users` and `posts` with a foreign key.
