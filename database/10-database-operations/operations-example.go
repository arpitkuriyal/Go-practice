// Package databaseoperations demonstrates context-aware queries and pool metrics.
package databaseoperations

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// FindUserName gives one database operation its own deadline. The caller's
// context remains the parent, so cancellation still flows through.
func FindUserName(ctx context.Context, db *sql.DB, id int64) (string, error) {
	queryContext, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var name string
	if err := db.QueryRowContext(queryContext,
		"SELECT name FROM users WHERE id = $1", id).Scan(&name); err != nil {
		return "", fmt.Errorf("find user %d: %w", id, err)
	}
	return name, nil
}

// ConfigurePool sets example limits. Choose real values from the database
// connection budget and observed workload rather than copying these blindly.
func ConfigurePool(db *sql.DB) {
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)
}

// PoolSnapshot contains the most useful pool gauges and counters for metrics.
// WaitCount and WaitDuration are cumulative counters, so monitoring systems
// should calculate their change over an interval.
type PoolSnapshot struct {
	MaxOpen      int
	Open         int
	InUse        int
	Idle         int
	WaitCount    int64
	WaitDuration time.Duration
}

func PoolStats(db *sql.DB) PoolSnapshot {
	stats := db.Stats()
	return PoolSnapshot{
		MaxOpen:      stats.MaxOpenConnections,
		Open:         stats.OpenConnections,
		InUse:        stats.InUse,
		Idle:         stats.Idle,
		WaitCount:    stats.WaitCount,
		WaitDuration: stats.WaitDuration,
	}
}
