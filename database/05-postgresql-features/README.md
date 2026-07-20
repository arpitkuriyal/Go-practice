# PostgreSQL Features: Revision

## Useful types and features

| Feature | Good use | Caution |
| --- | --- | --- |
| `UUID` | IDs exposed publicly or generated across services. | Larger than `BIGINT`; choose deliberately. |
| `JSONB` | Flexible attributes with occasional querying. | Do not turn a relational schema into one unvalidated document. |
| Arrays | Small, naturally multi-valued fields. | A join table is often better for relations that need constraints/querying. |
| `ON CONFLICT` | Idempotent create/upsert flows. | Define the conflict key and update semantics carefully. |
| Full-text search | Tokenized document search. | Different from substring search; index it appropriately. |

## Upsert

```sql
INSERT INTO users (email, name)
VALUES ($1, $2)
ON CONFLICT (email)
DO UPDATE SET name = EXCLUDED.name
RETURNING id, email, name;
```

An upsert is not automatically correct for every operation. Be explicit about which fields may be overwritten and who owns them.

## JSONB and indexing

```sql
CREATE TABLE events (
    id BIGSERIAL PRIMARY KEY,
    payload JSONB NOT NULL
);

CREATE INDEX events_payload_gin_idx ON events USING GIN (payload);
```

Use JSONB for metadata whose shape genuinely varies. Keep fields needed for joins, constraints, and frequent filters as ordinary columns whenever possible.
