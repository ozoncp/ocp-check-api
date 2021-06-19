package repo_test

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-check-api/internal/models"
	repo "github.com/ozoncp/ocp-check-api/internal/repo"
	"github.com/rs/zerolog"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

var _ = Describe("Repo", func() {
	const (
		errNotFound = "record not found"
	)

	var (
		mock      sqlmock.Sqlmock
		ctx       context.Context
		db        *sqlx.DB
		id        uint64
		err       error
		checkRepo repo.CheckRepo
		log       zerolog.Logger
	)

	BeforeEach(func() {
		db, mock, err = sqlmock.Newx() // mock sql.DB
		ctx = context.Background()
		log = zerolog.New(os.Stdout)
		checkRepo = repo.NewCheckRepo(db, &log, true)
	})

	AfterEach(func() {
		defer db.Close()
		err = mock.ExpectationsWereMet() // make sure all expectations were met
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("insert single check into database", func() {
		const (
			newId = uint64(100)
		)
		BeforeEach(func() {
			mock.ExpectBegin()
			rows := sqlmock.NewRows([]string{"id"}).AddRow(uint64(newId))
			mock.ExpectQuery("INSERT INTO checks").WithArgs(3, 4, 5, false).WillReturnRows(rows)
			mock.ExpectCommit()
			id, err = checkRepo.CreateCheck(ctx, models.Check{SolutionID: 3, TestID: 4, RunnerID: 5, Success: false})
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(newId))
		})
	})

	Context("insert multiple checks into database", func() {
		BeforeEach(func() {
			rows := sqlmock.NewRows([]string{"id"}).AddRow(uint64(1))
			rows2 := sqlmock.NewRows([]string{"id"}).AddRow(uint64(2))
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO checks").WithArgs(3, 4, 5, false).WillReturnRows(rows)
			mock.ExpectQuery("INSERT INTO checks").WithArgs(5, 6, 7, true).WillReturnRows(rows2)
			mock.ExpectCommit()
			_, err = checkRepo.MultiCreateCheck(ctx, []models.Check{
				{SolutionID: 3, TestID: 4, RunnerID: 5, Success: false},
				{SolutionID: 5, TestID: 6, RunnerID: 7, Success: true},
			})
		})

		It("", func() {
			Expect(err).Should(BeNil())
		})
	})

	Context("get check by id", func() {
		const (
			checkId = uint64(100)
		)
		var (
			check *models.Check
		)
		BeforeEach(func() {
			mockRows := sqlmock.NewRows([]string{"id", "solution_id", "test_id", "runner_id", "success"}).
				AddRow(checkId, 2, 3, 4, false)
			mock.ExpectQuery("SELECT id, solution_id, test_id, runner_id, success FROM checks").WithArgs(checkId).WillReturnRows(mockRows)
			check, err = checkRepo.DescribeCheck(ctx, checkId)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(check.ID).Should(Equal(checkId))
		})
	})

	Context("list checks", func() {
		var (
			limit  uint64
			offset uint64
			checks []models.Check
		)
		BeforeEach(func() {
			limit = 20
			offset = 1000
			mockRows := sqlmock.NewRows([]string{"id", "solution_id", "test_id", "runner_id", "success"}).
				AddRow(1, 2, 3, 4, false).AddRow(2, 3, 4, 5, true)
			mock.ExpectQuery(fmt.Sprintf("SELECT id, solution_id, test_id, runner_id, success FROM checks LIMIT %v OFFSET %v",
				limit, offset)).WillReturnRows(mockRows)
			checks, err = checkRepo.ListChecks(ctx, limit, offset)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(checks).Should(HaveLen(2))
		})
	})

	Context("remove unexisting check will result an error", func() {
		BeforeEach(func() {
			mock.ExpectExec("DELETE FROM checks").WithArgs(1).WillReturnError(errors.New(errNotFound))
			err = checkRepo.RemoveCheck(ctx, 1)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(Equal(errNotFound))
		})
	})
})
