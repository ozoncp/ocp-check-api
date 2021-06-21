// Package repo defines CheckRepo and TestRepo interfaces and implements types which
// provide basic CRUD operations for models.Check and models.Type
//
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
	// ErrCheckNotFound describes case when specified check not found in the database
	ErrCheckNotFound = errors.New("check not found")
	// ErrTestNotFound describes case when specified test not found in the database
	ErrTestNotFound = errors.New("test not found")
)

// CheckRepo interface, version 1
type CheckRepo interface {
	CreateCheck(ctx context.Context, check models.Check) (uint64, error)
	MultiCreateCheck(ctx context.Context, checks []models.Check) ([]uint64, error)
	UpdateCheck(ctx context.Context, check models.Check) (bool, error)
	RemoveCheck(ctx context.Context, checkId uint64) error
	DescribeCheck(ctx context.Context, checkId uint64) (*models.Check, error)
	ListChecks(ctx context.Context, limit, offset uint64) ([]models.Check, error)
}

// TestRepo interface, version 1
type TestRepo interface {
	CreateTest(ctx context.Context, test models.Test) (uint64, error)
	MultiCreateTest(ctx context.Context, tests []models.Test) ([]uint64, error)
	UpdateTest(ctx context.Context, test models.Test) (bool, error)
	RemoveTest(ctx context.Context, testId uint64) error
	DescribeTest(ctx context.Context, testId uint64) (*models.Test, error)
	ListTests(ctx context.Context, limit, offset uint64) ([]models.Test, error)
}

// CheckRepo implementation
type checkRepo struct {
	db  *sqlx.DB
	log *zerolog.Logger
}

// ListChecks returns database records (slice of models.Check) by offset. Parameter limit is a maximum size of slice.
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

// DescribeCheck searches record in database and returns a models.Check (pointer), if record found by checkId, ErrCheckNotFound
// if no such record found, or error in case of db error.
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
			return nil, ErrCheckNotFound
		default:
			r.log.Error().Err(err).Msg("")
			return nil, err
		}
	}

	return &check, nil
}

// UpdateCheck updates record in database and returns true if record was updated, ErrCheckNotFound
// if no such record found, or another error in case of db error.
func (r *checkRepo) UpdateCheck(ctx context.Context, check models.Check) (bool, error) {
	query := sq.Update("checks").
		Where(sq.Eq{"id": check.ID}).
		Set("solution_id", check.SolutionID).
		Set("test_id", check.TestID).
		Set("runner_id", check.RunnerID).
		Set("success", check.Success).
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
		return false, ErrCheckNotFound
	}

	return true, err
}

// RemoveCheck deletes record in database and returns ErrCheckNotFound
// if record was not found and deleted, or another error in case of db error.
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
		return ErrCheckNotFound
	}

	return nil
}

// CreateCheck creates new record in database and returns its id, or returns error in case of db error.
func (r *checkRepo) CreateCheck(ctx context.Context, check models.Check) (uint64, error) {
	sb := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("checks").
		Columns("solution_id", "test_id", "runner_id", "success").
		Values(check.SolutionID, check.TestID, check.RunnerID, check.Success).
		Suffix("RETURNING id")

	query, args, err := sb.ToSql()

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
		r.log.Debug().Err(err)
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	// Processing of rollback in case of error is not required
	return id, nil
}

// MultiCreateCheck creates batch of new records in database and returns their ids, or returns error in case of db error.
func (r *checkRepo) MultiCreateCheck(ctx context.Context, checks []models.Check) ([]uint64, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var ids []uint64

	for _, check := range checks {
		sb := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Insert("checks").
			Columns("solution_id", "test_id", "runner_id", "success").
			Suffix("RETURNING id")

		query, args, err := sb.
			Values(check.SolutionID, check.TestID, check.RunnerID, check.Success).
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

// NewCheckRepo creates new instance (checkRepo) of CheckRepo interface
func NewCheckRepo(db *sqlx.DB, log *zerolog.Logger) CheckRepo {
	return &checkRepo{db: db, log: log}
}

// TestRepo implementation
type testRepo struct {
	db  *sqlx.DB
	log *zerolog.Logger
}

// ListTests returns database records (slice of models.Test) by offset. Parameter limit is a maximum size of slice.
func (r *testRepo) ListTests(ctx context.Context, limit, offset uint64) ([]models.Test, error) {
	query := sq.Select("id", "task_id", "input", "output").
		From("tests").
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

	test := models.Test{}
	tests := []models.Test{}

	for rows.Next() {
		if err := rows.Scan(&test.ID, &test.TaskID, &test.Input, &test.Output); err != nil {
			r.log.Error().Err(err).Msg("")
			return nil, err
		}
		tests = append(tests, test)
	}

	return tests, nil
}

// DescribeTest searches record in database and returns a models.Test (pointer), if record found by testId, ErrCheckNotFound
// if no such record found, or error in case of db error.
func (r *testRepo) DescribeTest(ctx context.Context, testId uint64) (*models.Test, error) {
	query := sq.Select("id", "task_id", "input", "output").
		From("tests").
		Where(sq.Eq{"id": testId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)
	row := query.QueryRowContext(ctx)

	test := models.Test{}
	if err := row.Scan(&test.ID, &test.TaskID, &test.Input, &test.Output); err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrTestNotFound
		default:
			r.log.Error().Err(err).Msg("")
			return nil, err
		}
	}

	return &test, nil
}

// UpdateTest updates record in database and returns true if record was updated, ErrTestNotFound
// if no such record found, or another error in case of db error.
func (r *testRepo) UpdateTest(ctx context.Context, test models.Test) (bool, error) {
	query := sq.Update("tests").
		Where(sq.Eq{"id": test.ID}).
		Set("task_id", test.TaskID).
		Set("input", test.Input).
		Set("output", test.Output).
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
		return false, ErrTestNotFound
	}

	return true, err
}

// RemoveTest deletes record in database and returns ErrTestNotFound
// if record was not found and deleted, or another error in case of db error.
func (r *testRepo) RemoveTest(ctx context.Context, testId uint64) error {
	query := sq.Delete("tests").
		Where(sq.Eq{"id": testId}).
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
		return ErrTestNotFound
	}

	return nil
}

// CreateTest creates new record in database and returns its id, or returns error in case of db error.
func (r *testRepo) CreateTest(ctx context.Context, test models.Test) (uint64, error) {
	sb := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("tests").
		Columns("task_id", "input", "output").
		Values(test.TaskID, test.Input, test.Output).
		Suffix("RETURNING id")

	query, args, err := sb.ToSql()

	r.log.Debug().Msgf("%v", query)

	if err != nil {
		return 0, err
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}

	var id uint64
	if err := tx.QueryRowxContext(ctx, query, args[0], args[1], args[2]).Scan(&id); err != nil {
		r.log.Debug().Err(err)
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	// Processing of rollback in case of error is not required
	return id, nil
}

// MultiCreateTest creates batch of new records in database and returns their ids, or returns error in case of db error.
func (r *testRepo) MultiCreateTest(ctx context.Context, tests []models.Test) ([]uint64, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var ids []uint64

	for _, test := range tests {
		sb := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Insert("tests").
			Columns("task_id", "input", "output").
			Suffix("RETURNING id")

		query, args, err := sb.
			Values(test.TaskID, test.Input, test.Output).
			ToSql()

		if err != nil {
			return nil, err
		}

		var id uint64
		if err := tx.QueryRowxContext(ctx, query, args[0], args[1], args[2]).Scan(&id); err != nil {
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

// NewTestRepo creates new instance (testRepo) of TestRepo interface
func NewTestRepo(db *sqlx.DB, log *zerolog.Logger) TestRepo {
	return &testRepo{db: db, log: log}
}
