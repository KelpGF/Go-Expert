package usecase

import (
	"context"

	"github.com/KelpGF/Go-Expert/017-UOW/internal/entity"
	"github.com/KelpGF/Go-Expert/017-UOW/internal/repository"
	"github.com/KelpGF/Go-Expert/017-UOW/pkg/uow"
)

type AddCourseUseCaseUow struct {
	Uow uow.UnitOfWork
}

func NewAddCourseUseCaseUow(uow uow.UnitOfWork) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(r uow.UnitOfWork) error {
		category := entity.Category{
			ID:   input.CourseCategoryID,
			Name: input.CategoryName,
		}

		CategoryRepository, err := a.getCategoryRepository(ctx)
		if err != nil {
			return err
		}

		CourseRepository, err := a.getCourseRepository(ctx)
		if err != nil {
			return err
		}

		err = CategoryRepository.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		err = CourseRepository.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) (repository.CategoryRepositoryInterface, error) {
	CategoryRepository, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		return nil, err
	}

	return CategoryRepository.(repository.CategoryRepositoryInterface), nil
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) (repository.CourseRepositoryInterface, error) {
	CourseRepository, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		return nil, err
	}

	return CourseRepository.(repository.CourseRepositoryInterface), nil
}
