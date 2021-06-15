package repo

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/ozoncp/ocp-check-api/internal/models"
	"github.com/rs/zerolog"
)

var (
	CheckNotFound = errors.New("check not found")
)

type CheckRepo interface {
	AddCheck(check models.Check) (uint64, error)
	AddChecks(checks []models.Check) error
	RemoveCheck(checkId uint64) error
	DescribeCheck(checkId uint64) (*models.Check, error)
	ListChecks(limit, offset uint64) ([]models.Check, error)
}

type TestRepo interface {
	AddTests(tests []models.Test) error
	RemoveTest(testId uint64) error
	DescribeTest(testId uint64) (*models.Test, error)
	ListTests(limit, offset uint64) ([]models.Test, error)
}

type checkRepo struct {
	db  *sql.DB
	ctx *context.Context
	log *zerolog.Logger
}

func (r *checkRepo) ListChecks(limit, offset uint64) ([]models.Check, error) {
	query := sq.Select("id", "solution_id", "test_id", "runner_id", "success").
		From("checks").
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)
	rows, err := query.QueryContext(*r.ctx)
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

func (r *checkRepo) DescribeCheck(checkId uint64) (*models.Check, error) {
	query := sq.Select("id", "solution_id", "test_id", "runner_id", "success").
		From("checks").
		Where(sq.Eq{"id": checkId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)
	row := query.QueryRowContext(*r.ctx)

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

func (r *checkRepo) RemoveCheck(checkId uint64) error {
	query := sq.Delete("checks").
		Where(sq.Eq{"id": checkId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var result sql.Result
	result, err := query.ExecContext(*r.ctx)
	if err != nil {
		r.log.Error().Err(err).Msg("")
	}

	// no rows affected and no error, it is a case of record not found
	rows, resultErr := result.RowsAffected()
	if rows == int64(0) && resultErr == nil {
		return CheckNotFound
	}

	return err
}

func (r *checkRepo) AddCheck(check models.Check) (uint64, error) {
	query := sq.Insert("checks").
		Columns("solution_id", "test_id", "runner_id", "success").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	query = query.Values(check.SolutionID, check.TestID, check.RunnerID, check.Success)

	result, err := query.ExecContext(*r.ctx)
	if err != nil {
		r.log.Error().Err(err).Msg("")
	}

	id, _ := result.LastInsertId()
	return uint64(id), err
}

func (r *checkRepo) AddChecks(checks []models.Check) error {
	query := sq.Insert("checks").
		Columns("solution_id", "test_id", "runner_id", "success").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, check := range checks {
		query = query.Values(check.SolutionID, check.TestID, check.RunnerID, check.Success)
	}

	_, err := query.ExecContext(*r.ctx)
	if err != nil {
		r.log.Error().Err(err).Msg("")
	}

	return err
}

func NewCheckRepo(ctx *context.Context, db *sql.DB, log *zerolog.Logger) CheckRepo {
	return &checkRepo{db: db, ctx: ctx, log: log}
}
