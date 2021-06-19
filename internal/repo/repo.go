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
	TestNotFound  = errors.New("test not found")
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
	db       *sqlx.DB
	log      *zerolog.Logger
	forTests bool
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
	sb := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("checks").
		Columns("solution_id", "test_id", "runner_id", "success").
		Values(check.SolutionID, check.TestID, check.RunnerID, check.Success)
	if !r.forTests {
		sb = sb.Suffix("RETURNING id")
	}
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
			Columns("solution_id", "test_id", "runner_id", "success")
		if !r.forTests {
			sb = sb.Suffix("RETURNING id")
		}

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

func NewCheckRepo(db *sqlx.DB, log *zerolog.Logger, forTests bool) CheckRepo {
	return &checkRepo{db: db, log: log, forTests: forTests}
}

// Test REPO

type testRepo struct {
	db       *sqlx.DB
	log      *zerolog.Logger
	forTests bool
}

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

func (r *testRepo) DescribeTest(ctx context.Context, checkId uint64) (*models.Test, error) {
	query := sq.Select("id", "task_id", "input", "output").
		From("tests").
		Where(sq.Eq{"id": checkId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)
	row := query.QueryRowContext(ctx)

	test := models.Test{}
	if err := row.Scan(&test.ID, &test.TaskID, &test.Input, &test.Output); err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, TestNotFound
		default:
			r.log.Error().Err(err).Msg("")
			return nil, err
		}
	}

	return &test, nil
}

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
		return false, TestNotFound
	}

	return true, err
}

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
		return TestNotFound
	}

	return nil
}

func (r *testRepo) CreateTest(ctx context.Context, test models.Test) (uint64, error) {
	sb := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("tests").
		Columns("task_id", "input", "output").
		Values(test.TaskID, test.Input, test.Output)
	if !r.forTests {
		sb = sb.Suffix("RETURNING id")
	}
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
			Columns("task_id", "input", "output")
		if !r.forTests {
			sb = sb.Suffix("RETURNING id")
		}

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

func NewTestRepo(db *sqlx.DB, log *zerolog.Logger, forTests bool) TestRepo {
	return &testRepo{db: db, log: log, forTests: forTests}
}
