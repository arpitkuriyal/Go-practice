-- A simplified PostgreSQL money-transfer flow.
-- Run all statements in one transaction: either both balances change or neither does.

BEGIN;

-- Lock both account rows before checking/updating their balances.
SELECT id, balance
FROM accounts
WHERE id IN (1, 2)
ORDER BY id
FOR UPDATE;

-- After checking that account 1 has enough money:
UPDATE accounts SET balance = balance - 500 WHERE id = 1;
UPDATE accounts SET balance = balance + 500 WHERE id = 2;

COMMIT;
-- Use ROLLBACK instead of COMMIT if any step fails.
