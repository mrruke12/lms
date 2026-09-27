package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mrruke12/lms/internal/domains/revision"
)

type RevisionRepository struct {
	pool *pgxpool.Pool
}

func (r *RevisionRepository) GetByLessonID(ctx context.Context, id uuid.UUID) ([]revision.Revision, error) {
	sql := `
		select distinct on (element_id, lesson_id)
			element_id, id, lesson_id, parent_id, type, assessment, config
		from revisions
		where lesson_id = $1
		order by element_id, lesson_id, created_at desc
	`

	rows, err := r.pool.Query(ctx, sql, id)

	if err != nil {
		return nil, wrapDriverErr(err)
	}

	defer rows.Close()

	var revs []revision.Revision

	for rows.Next() {
		rev, err := revision.FromRow(rows)

		if err != nil {
			return nil, wrapDriverErr(err)
		}

		revs = append(revs, *rev)
	}

	if err = rows.Err(); err != nil {
		return nil, wrapDriverErr(err)
	}

	return revs, nil
}

func (r *RevisionRepository) BulkCreate(ctx context.Context, revs []revision.Revision) error {
	query := builder.Insert("elements").Columns("element_id", "id", "lesson_id", "parent_id", "type", "assessment", "config")

	for _, rev := range revs {
		query = query.Values(rev.ElementID(), rev.ID(), rev.LessonID(), rev.ParentID(), rev.Type(), rev.Assessment(), rev.ConfigRaw())
	}

	sql, args, err := query.ToSql()

	if err != nil {
		return wrapDriverErr(err)
	}

	_, err = r.pool.Exec(ctx, sql, args...)

	return wrapDriverErr(err)
}
