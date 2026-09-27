package postgres

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mrruke12/lms/internal/domains/domainerr"
	"github.com/mrruke12/lms/internal/domains/element"
)

type ElementRepository struct {
	pool *pgxpool.Pool
}

func (r *ElementRepository) GetByID(ctx context.Context, id uuid.UUID) (*element.Element, error) {
	sql := `
		select id, lesson_id, parent_id, typ, assessment, config
		from elements
		where id = $1
	`

	row := r.pool.QueryRow(ctx, sql, id)
	el, err := element.FromRow(row)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domainerr.ErrNotFound
		}

		return nil, err
	}

	return el, nil
}

func (r *ElementRepository) GetByLessonID(ctx context.Context, id uuid.UUID) ([]element.Element, error) {
	sql, args, err := builder.Select("id", "lesson_id", "parent_id", "type", "assessment", "config").
		From("elements").
		Where(squirrel.Eq{"lesson_id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var els []element.Element

	for rows.Next() {
		el, err := element.FromRow(rows)

		if err != nil {
			return nil, err
		}

		els = append(els, *el)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return els, nil
}

func (r *ElementRepository) BulkCreate(ctx context.Context, els []element.Element) error {
	query := builder.Insert("elements").Columns("id", "lesson_id", "parent_id", "type", "assessment", "config")

	for _, el := range els {
		query = query.Values(el.ID(), el.LessonID(), el.ParentID(), el.Type(), el.Assessment(), el.ConfigRaw())
	}

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.pool.Exec(ctx, sql, args...)

	return err
}

func (r *ElementRepository) BulkUpdate(ctx context.Context, els []element.Element) error {
	l := len(els)

	ids := make([]uuid.UUID, l)
	parentIDs := make([]*uuid.UUID, l)
	configs := make([]json.RawMessage, l)

	for i, el := range els {
		ids[i] = el.ID()
		parentIDs[i] = el.ParentID()
		configs[i] = el.ConfigRaw()
	}

	sql := `
		update elements as e
		set
			parent_id = v.parent_id,
			config = v.config
		from unnest(
			$1::uuid[],
			$2::uuid[],
			$3::jsonb[]
		) as v(id, parent_id, config)
		where e.id = v.id
	`

	_, err := r.pool.Exec(
		ctx,
		sql,
		ids, parentIDs, configs,
	)

	return err
}

func (r *ElementRepository) BulkDelete(ctx context.Context, ids []uuid.UUID) error {
	sql := `
		delete from elements
		where id = any($1::uuid[]) 
	`

	_, err := r.pool.Exec(
		ctx,
		sql,
		ids,
	)

	return err
}
