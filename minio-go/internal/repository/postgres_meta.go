package repository

import (
	"context"
	"database/sql"
)

type PostgresMetaRepository struct {
	db *sql.DB
}

func NewPostgresMetaRepository(db *sql.DB) *PostgresMetaRepository {
	return &PostgresMetaRepository{db: db}
}

func (r *PostgresMetaRepository) GetByIdempotencyKey(ctx context.Context, key string) (*PasteMeta, error) {
	const q = `
		SELECT id, idempotency_key, filename, created_at, expires_at 
		FROM paste_metadata 
		WHERE idempotency_key = $1;
	`
	paste := &PasteMeta{}

	err := r.db.QueryRowContext(ctx, q, key).Scan(
		&paste.ID,
		&paste.IdempotencyKey,
		&paste.Filename,
		&paste.CreatedAt,
		&paste.ExpiresAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return paste, nil
}

func (r *PostgresMetaRepository) GetByID(ctx context.Context, id string) (*PasteMeta, error) {
	const q = `
		SELECT id, idempotency_key, filename, created_at, expires_at 
		FROM paste_metadata 
		WHERE id = $1;
	`
	paste := &PasteMeta{}

	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&paste.ID,
		&paste.IdempotencyKey,
		&paste.Filename,
		&paste.CreatedAt,
		&paste.ExpiresAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return paste, nil

}

func (r *PostgresMetaRepository) Insert(ctx context.Context, meta *PasteMeta) (*PasteMeta, error) {
	const q = `
        INSERT INTO paste_metadata (
			id,
            idempotency_key,
            filename,
            created_at,
            expires_at
        )
        VALUES ($1, $2, $3, $4, $5) RETURNING id;
    `
	inserted := &PasteMeta{}

	err := r.db.QueryRowContext(ctx, q,
		meta.ID,
		meta.IdempotencyKey,
		meta.Filename,
		meta.CreatedAt,
		meta.ExpiresAt,
	).Scan(
		&inserted.ID,
	)

	if err != nil {
		return nil, err
	}

	return inserted, nil
}
