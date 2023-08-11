package sqldb_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dewzzjr/ais/internal/model"
	"github.com/dewzzjr/ais/internal/repository/sqldb"
	"github.com/dewzzjr/ais/pkg/pointer"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setup(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	t.Helper()
	mockdb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		Conn:                      mockdb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, err)
	return gormDB, mock
}

func TestFetchArticles(t *testing.T) {
	t.Run("ShouldReturnResult_WhenExpectSelectAll", func(t *testing.T) {
		db, mock := setup(t)
		mock.ExpectQuery(
			"SELECT (.+) FROM `articles` WHERE `articles`.`deleted_at` IS NULL ORDER BY `created_at` DESC").
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "title", "body", "author", "created_at", "updated_at"}).
					AddRow(1, "title test", "body test", "author test", time.Now().Truncate(time.Minute), sql.NullTime{}),
			)
		expectResult := []model.Article{
			{Author: "author test", Title: "title test", Body: "body test", Model: gorm.Model{ID: 1, CreatedAt: time.Now().Truncate(time.Minute)}},
		}
		r := sqldb.New(db)
		gotResult, err := r.FetchArticles(context.Background(), model.Filter{})
		assert.NoError(t, err)
		assert.EqualValues(t, expectResult, gotResult)
	})
	t.Run("ShouldReturnResult_WhenExpectSelectWithFilter", func(t *testing.T) {
		db, mock := setup(t)
		mock.ExpectQuery(
			"SELECT (.+) FROM `articles` WHERE MATCH(.+) AGAINST(.+) AND author = (.+) AND `articles`.`deleted_at` IS NULL ORDER BY `created_at` DESC").
			WithArgs("test", "author test").
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "title", "body", "author", "created_at", "updated_at"}).
					AddRow(1, "title test", "body test", "author test", time.Now().Truncate(time.Minute), sql.NullTime{}),
			)
		expectResult := []model.Article{
			{Author: "author test", Title: "title test", Body: "body test", Model: gorm.Model{ID: 1, CreatedAt: time.Now().Truncate(time.Minute)}},
		}
		r := sqldb.New(db)
		gotResult, err := r.FetchArticles(context.Background(), model.Filter{Query: pointer.New("test"), Author: pointer.New("author test")})
		assert.NoError(t, err)
		assert.EqualValues(t, expectResult, gotResult)
	})
}

func TestInsertArticles(t *testing.T) {
	t.Run("ShouldNoError_WhenExpectInsertSuccess", func(t *testing.T) {
		db, mock := setup(t)
		mock.ExpectBegin()
		mock.ExpectExec(
			"INSERT INTO `articles` (.+) VALUES (.+)").
			WillReturnResult(sqlmock.NewResult(10, 1))
		mock.ExpectCommit()
		payload := model.Article{Author: "author test", Title: "title test", Body: "body test"}
		r := sqldb.New(db)
		got, err := r.InsertArticle(context.Background(), payload)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.EqualValues(t, 10, got.ID)
	})
}
