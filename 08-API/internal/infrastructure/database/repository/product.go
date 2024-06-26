package repository

import (
	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (productRepository *ProductRepository) Create(product *entity.Product) error {
	return productRepository.DB.Create(product).Error
}

func (productRepository *ProductRepository) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	var products []*entity.Product

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	if sort == "" {
		sort = "asc"
	}

	offset := (page - 1) * limit
	sortColumn := "created_at " + sort

	err := productRepository.DB.Offset(offset).Limit(limit).Order(sortColumn).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (productRepository *ProductRepository) FindByID(id string) (*entity.Product, error) {
	var product entity.Product

	err := productRepository.DB.First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (productRepository *ProductRepository) Update(product *entity.Product) error {
	_, err := productRepository.FindByID(product.ID.String())
	if err != nil {
		return err
	}

	return productRepository.DB.Save(product).Error
}

func (productRepository *ProductRepository) Delete(id string) error {
	_, err := productRepository.FindByID(id)
	if err != nil {
		return err
	}

	return productRepository.DB.Where("id = ?", id).Delete(&entity.Product{}).Error
}
