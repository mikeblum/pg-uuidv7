// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: batch.go

package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrBatchAlreadyClosed = errors.New("batch already closed")
)

const insertUUIDv4Bulk = `-- name: InsertUUIDv4Bulk :batchmany

INSERT INTO uuid_v4(id, created)
VALUES($1, $2)
`

type InsertUUIDv4BulkBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type InsertUUIDv4BulkParams struct {
	ID      pgtype.UUID
	Created pgtype.Timestamp
}

type InsertUUIDv4BulkRow struct {
}

func (q *Queries) InsertUUIDv4Bulk(ctx context.Context, arg []InsertUUIDv4BulkParams) *InsertUUIDv4BulkBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.ID,
			a.Created,
		}
		batch.Queue(insertUUIDv4Bulk, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &InsertUUIDv4BulkBatchResults{br, len(arg), false}
}

func (b *InsertUUIDv4BulkBatchResults) Query(f func(int, []InsertUUIDv4BulkRow, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		var items []InsertUUIDv4BulkRow
		if b.closed {
			if f != nil {
				f(t, items, ErrBatchAlreadyClosed)
			}
			continue
		}
		err := func() error {
			rows, err := b.br.Query()
			if err != nil {
				return err
			}
			defer rows.Close()
			for rows.Next() {
				var i InsertUUIDv4BulkRow
				if err := rows.Scan(); err != nil {
					return err
				}
				items = append(items, i)
			}
			return rows.Err()
		}()
		if f != nil {
			f(t, items, err)
		}
	}
}

func (b *InsertUUIDv4BulkBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const insertUUIDv7Bulk = `-- name: InsertUUIDv7Bulk :batchmany

INSERT INTO uuid_v7(id, created)
VALUES($1, $2)
`

type InsertUUIDv7BulkBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type InsertUUIDv7BulkParams struct {
	ID      pgtype.UUID
	Created pgtype.Timestamp
}

type InsertUUIDv7BulkRow struct {
}

func (q *Queries) InsertUUIDv7Bulk(ctx context.Context, arg []InsertUUIDv7BulkParams) *InsertUUIDv7BulkBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.ID,
			a.Created,
		}
		batch.Queue(insertUUIDv7Bulk, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &InsertUUIDv7BulkBatchResults{br, len(arg), false}
}

func (b *InsertUUIDv7BulkBatchResults) Query(f func(int, []InsertUUIDv7BulkRow, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		var items []InsertUUIDv7BulkRow
		if b.closed {
			if f != nil {
				f(t, items, ErrBatchAlreadyClosed)
			}
			continue
		}
		err := func() error {
			rows, err := b.br.Query()
			if err != nil {
				return err
			}
			defer rows.Close()
			for rows.Next() {
				var i InsertUUIDv7BulkRow
				if err := rows.Scan(); err != nil {
					return err
				}
				items = append(items, i)
			}
			return rows.Err()
		}()
		if f != nil {
			f(t, items, err)
		}
	}
}

func (b *InsertUUIDv7BulkBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}
