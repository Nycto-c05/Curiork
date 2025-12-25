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
		SELECT uuid, idempotency_key, filename, created_at, expires_at 
		FROM paste_metadata 
		WHERE idempotency_key = $1;
	`
	paste := &PasteMeta{}

	err := r.db.QueryRowContext(ctx, q, key).Scan(
		&paste.UUID,
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
func (r *PostgresMetaRepository) GetbyUUID(ctx context.Context, uuid string) (*PasteMeta, error) {
	const q = `
		SELECT uuid, idempotency_key, filename, created_at, expires_at 
		FROM paste_metadata 
		WHERE uuid = $1;
	`
	paste := &PasteMeta{}

	err := r.db.QueryRowContext(ctx, q, uuid).Scan(
		&paste.UUID,
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
func (r *PostgresMetaRepository) Insert(ctx context.Context, meta *PasteMeta) error {
	const q = `
        INSERT INTO paste_metadata (
            uuid,
            idempotency_key,
            filename,
            created_at,
            expires_at
        )
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := r.db.ExecContext(ctx, q,
		meta.UUID,
		meta.IdempotencyKey,
		meta.Filename,
		meta.CreatedAt,
		meta.ExpiresAt,
	)

	return err

}
