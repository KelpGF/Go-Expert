package repository

import (
	"context"
	"database/sql"

	"github.com/KelpGF/Go-Expert/017-UOW/internal/db"
	"github.com/KelpGF/Go-Expert/017-UOW/internal/entity"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(dtb *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return r.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:   category.ID,
		Name: category.Name,
	})
}
