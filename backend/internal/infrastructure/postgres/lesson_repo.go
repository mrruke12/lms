package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mrruke12/lms/internal/domains/domainerr"
	"github.com/mrruke12/lms/internal/domains/lesson"
)

type LessonRepository struct {
	pool *pgxpool.Pool
}

func NewLessonRepository(pool *pgxpool.Pool) *LessonRepository {
	return &LessonRepository{
		pool: pool,
	}
}

func (r *LessonRepository) GetByID(ctx context.Context, id uuid.UUID) (*lesson.Lesson, error) {
	sql := `
		select id, author_id, version, name, status
		from lessons
		where id = $1
	`

	row := r.pool.QueryRow(ctx, sql, id)
	l, err := lesson.FromRow(row)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domainerr.ErrNotFound
		}

		return nil, err
	}

	return l, nil
}

func (r *LessonRepository) Create(ctx context.Context, l *lesson.Lesson) error {
	sql := `
		insert into lessons (id, author_id, version, name, status)
		values ($1, $2, $3, $4, $5)
	`

	_, err := r.pool.Exec(
		ctx,
		sql,
		l.ID(), l.AuthorID(), l.Version(), l.Name(), l.Status(),
	)

	return err
}

func (r *LessonRepository) Update(ctx context.Context, l *lesson.Lesson) error {
	sql := `
		update lessons
		set version = $1,
		name = $2,
		status = $3
		where id = $4
	`

	tag, err := r.pool.Exec(
		ctx,
		sql,
		l.Version(), l.Name(), l.Status(), l.ID(),
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return domainerr.ErrNotFound
	}

	return nil
}
