package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/KelpGF/Go-Expert/017-UOW/internal/db"
	"github.com/KelpGF/Go-Expert/017-UOW/internal/repository"
	"github.com/KelpGF/Go-Expert/017-UOW/pkg/uow"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourseUow(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(go_uow_mysql:3306)/courses")
	assert.NoError(t, err)

	uow, _ := uow.NewUow(dbt)
	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbt)
		repo.Queries = db.New(tx)

		return repo
	})
	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbt)
		repo.Queries = db.New(tx)

		return repo
	})

	input := InputUseCase{
		CategoryName:     "Category Uow 1",
		CourseName:       "Course Uow 1",
		CourseCategoryID: uuid.New().String(),
	}

	ctx := context.Background()

	useCase := NewAddCourseUseCaseUow(uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
