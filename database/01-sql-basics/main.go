// Package sqlbasics shows the database/sql patterns used by a small repository.
package sqlbasics

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID    int64
	Name  string
	Email string
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

// CreateUser uses placeholders rather than interpolating values into SQL.
// The $1 syntax is for PostgreSQL; placeholder syntax depends on the driver.
func (r Repository) CreateUser(ctx context.Context, user User) (int64, error) {
	const query = `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	if err := r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&user.ID); err != nil {
		return 0, fmt.Errorf("create user: %w", err)
	}
	return user.ID, nil
}

func (r Repository) FindUser(ctx context.Context, id int64) (User, error) {
	const query = `SELECT id, name, email FROM users WHERE id = $1`
	var user User
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user %d: %w", id, sql.ErrNoRows)
		}
		return User{}, fmt.Errorf("find user %d: %w", id, err)
	}
	return user, nil
}

// Transfer demonstrates that related writes must share one transaction.
func (r Repository) Transfer(ctx context.Context, fromID, toID, cents int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transfer: %w", err)
	}
	defer tx.Rollback() // safe no-op after Commit

	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance = balance - $1 WHERE id = $2`, cents, fromID); err != nil {
		return fmt.Errorf("debit account: %w", err)
	}
	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance = balance + $1 WHERE id = $2`, cents, toID); err != nil {
		return fmt.Errorf("credit account: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transfer: %w", err)
	}
	return nil
}

// ConfigurePool is an example starting point. Tune with production metrics.
func ConfigurePool(db *sql.DB) {
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)
}
