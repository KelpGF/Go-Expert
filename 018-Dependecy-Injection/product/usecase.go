package product

type ProductUseCase struct {
	repository ProductRepository
}

func NewProductUseCase(repository ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository}
}

func (u *ProductUseCase) GetProduct(id int) (*Product, error) {
	return u.repository.GetProduct(id)
}
