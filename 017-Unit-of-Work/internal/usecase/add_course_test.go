package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/KelpGF/Go-Expert/017-UOW/internal/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(go_uow_mysql:3306)/courses")
	assert.NoError(t, err)

	input := InputUseCase{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: uuid.New().String(),
	}

	ctx := context.Background()

	useCase := NewAddCourseUseCase(repository.NewCourseRepository(dbt), repository.NewCategoryRepository(dbt))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
