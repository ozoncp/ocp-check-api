package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
	"github.com/ozoncp/ocp-check-api/internal/models"
	"github.com/rs/zerolog"
)

var (
	CheckNotFound = errors.New("check not found")
)

type CheckRepo interface {
	CreateCheck(ctx context.Context, check models.Check) (uint64, error)
	MultiCreateCheck(ctx context.Context, checks []models.Check) ([]uint64, error)
	UpdateCheck(ctx context.Context, check models.Check) (bool, error)
	RemoveCheck(ctx context.Context, checkId uint64) error
	DescribeCheck(ctx context.Context, checkId uint64) (*models.Check, error)
	ListChecks(ctx context.Context, limit, offset uint64) ([]models.Check, error)
}

type TestRepo interface {
	CreateTest(ctx context.Context, test models.Test) (uint64, error)
	MultiCreateTest(ctx context.Context, tests []models.Test) ([]uint64, error)
	UpdateTest(ctx context.Context, test models.Test) (bool, error)
	RemoveTest(ctx context.Context, testId uint64) error
	DescribeTest(ctx context.Context, testId uint64) (*models.Test, error)
	ListTests(ctx context.Context, limit, offset uint64) ([]models.Test, error)
}

type checkRepo struct {
	db  *sqlx.DB
	log *zerolog.Logger
}

func (r *checkRepo) ListChecks(ctx context.Context, limit, offset uint64) ([]models.Check, error) {
	query := sq.Select("id", "solution_id", "test_id", "runner_id", "success").
		From("checks").
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)
	rows, err := query.QueryContext(ctx)
	if err != nil {
		r.log.Error().Err(err).Msg("")
		return nil, err
	}
	defer rows.Close()

	check := models.Check{}
	checks := []models.Check{}

	for rows.Next() {
		if err := rows.Scan(&check.ID, &check.SolutionID, &check.TestID, &check.RunnerID, &check.Success); err != nil {
			r.log.Error().Err(err).Msg("")
			return nil, err
		}
		checks = append(checks, check)
	}

	return checks, nil
}

func (r *checkRepo) DescribeCheck(ctx context.Context, checkId uint64) (*models.Check, error) {
	query := sq.Select("id", "solution_id", "test_id", "runner_id", "success").
		From("checks").
		Where(sq.Eq{"id": checkId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)
	row := query.QueryRowContext(ctx)

	check := models.Check{}
	if err := row.Scan(&check.ID, &check.SolutionID, &check.TestID, &check.RunnerID, &check.Success); err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, CheckNotFound
		default:
			r.log.Error().Err(err).Msg("")
			return nil, err
		}
	}

	return &check, nil
}

func (r *checkRepo) UpdateCheck(ctx context.Context, check models.Check) (bool, error) {
	query := sq.Update("checks").
		Where(sq.Eq{"id": check.ID}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var result sql.Result
	result, err := query.ExecContext(ctx)
	if err != nil {
		r.log.Error().Err(err).Msg("")
	}

	// no rows affected and no error, it is a case of record not found
	rows, resultErr := result.RowsAffected()
	if rows == int64(0) && resultErr == nil {
		return false, CheckNotFound
	}

	return true, err
}

func (r *checkRepo) RemoveCheck(ctx context.Context, checkId uint64) error {
	query := sq.Delete("checks").
		Where(sq.Eq{"id": checkId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var result sql.Result
	result, err := query.ExecContext(ctx)
	if err != nil {
		r.log.Error().Err(err).Msg("")
		return err
	}

	// no rows affected and no error, it is a case of record not found
	rows, resultErr := result.RowsAffected()
	if rows == int64(0) && resultErr == nil {
		return CheckNotFound
	}

	return nil
}

func (r *checkRepo) CreateCheck(ctx context.Context, check models.Check) (uint64, error) {
	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("checks").
		Columns("solution_id", "test_id", "runner_id", "success").
		Values(check.SolutionID, check.TestID, check.RunnerID, check.Success).
		Suffix("RETURNING id").
		ToSql()

	r.log.Debug().Msgf("%v", query)

	if err != nil {
		return 0, err
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}

	var id uint64
	if err := tx.QueryRowxContext(ctx, query, args[0], args[1], args[2], args[3]).Scan(&id); err != nil {
		return 0, err
	}

	tx.Commit()
	// Processing of rollback in case of error is not required
	return id, nil
}

func (r *checkRepo) MultiCreateCheck(ctx context.Context, checks []models.Check) ([]uint64, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var ids []uint64

	for _, check := range checks {
		query, args, err := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Insert("checks").
			Columns("solution_id", "test_id", "runner_id", "success").
			Values(check.SolutionID, check.TestID, check.RunnerID, check.Success).
			Suffix("RETURNING id").
			ToSql()

		if err != nil {
			return nil, err
		}

		var id uint64
		if err := tx.QueryRowxContext(ctx, query, args[0], args[1], args[2], args[3]).Scan(&id); err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	// Processing of rollback in case of error is not required
	return ids, nil
}

func NewCheckRepo(db *sqlx.DB, log *zerolog.Logger) CheckRepo {
	return &checkRepo{db: db, log: log}
}
