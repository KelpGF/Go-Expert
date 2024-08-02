package service

import (
	"context"
	"io"

	"github.com/KelpGF/Go-gRPC/internal/database"
	"github.com/KelpGF/Go-gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer

	categoryDb *database.Category
}

func NewCategoryService(categoryDb *database.Category) *CategoryService {
	return &CategoryService{categoryDb: categoryDb}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := s.categoryDb.Create(req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	res := s.parseCategory(category)

	return res, nil
}

func (s *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		createdCategory, err := s.categoryDb.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, s.parseCategory(createdCategory))
	}
}

func (s *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		createdCategory, err := s.categoryDb.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		err = stream.Send(s.parseCategory(createdCategory))
		if err != nil {
			return err
		}
	}
}

func (s *CategoryService) GetCategoryById(ctx context.Context, req *pb.GetCategoryByIdRequest) (*pb.Category, error) {
	category, err := s.categoryDb.FindByID(req.Id)
	if err != nil {
		return nil, err
	}

	res := s.parseCategory(category)

	return res, nil
}

func (s *CategoryService) ListCategories(ctx context.Context, req *pb.Blank) (*pb.CategoryList, error) {
	categories, err := s.categoryDb.FindAll()
	if err != nil {
		return nil, err
	}

	res := s.makeListResponse(categories)

	return res, nil
}

func (s *CategoryService) makeListResponse(categories []database.Category) *pb.CategoryList {
	var response []*pb.Category

	for _, category := range categories {
		response = append(response, s.parseCategory(&category))
	}

	return &pb.CategoryList{Categories: response}
}

func (s *CategoryService) parseCategory(category *database.Category) *pb.Category {
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}
