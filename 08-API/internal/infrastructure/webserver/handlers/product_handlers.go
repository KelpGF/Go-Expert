package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/dto"
	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"
	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/repository"
	pkgEntity "github.com/KelpGF/Go-Expert/08-APIs/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	repository repository.ProductRepository
}

func NewProductHandler(repository repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repository: repository}
}

// Create Product godoc
// @Summary 		Create a new product
// @Description Create a new product
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			request	body dto.CreateProductInput true "Product Request"
// @Success 		201 {object} entity.Product
// @Failure 		400 {object} Error
// @Failure 		500 {object} Error
// @Router 			/product [post]
// @Security  ApiKeyAuth
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse := Error{
			Message: "Invalid request body",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	product, err := entity.NewProduct(input.Name, input.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	err = h.repository.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := Error{
			Message: "Internal server error",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusCreated)
}

// Get Product godoc
// @Summary 		Get product
// @Description Get product a product by id
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			id	path string	true	"Product ID" format(uuid)
// @Success 		200 {object} entity.Product
// @Failure 		404 {object} Error
// @Failure 		500 {object} Error
// @Router 			/product/{id} [get]
// @Security  ApiKeyAuth
func (h *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	product, err := h.repository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errResponse := Error{
			Message: "Product not found",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// GetByPagination godoc
// @Summary 		Get products by pagination
// @Description Get products by pagination
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			page	query	int	false	"Page Number"
// @Param 			limit	query	int	false	"Limit of products"
// @Param 			enumstring sort	query	string	false	"Sort by" Enums(asc, desc)
// @Success 		200 {array} entity.Product
// @Failure 		500 {object} Error
// @Router 			/product [get]
// @Security  ApiKeyAuth
func (h *ProductHandler) GetByPagination(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	products, err := h.repository.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := Error{
			Message: "Internal server error",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	json.NewEncoder(w).Encode(products)
	w.WriteHeader(http.StatusOK)
}

// Edit Product godoc
// @Summary 		Edit product
// @Description Edit product a product by id
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			id	path string	true	"Product ID" format(uuid)
// @Param 			request	body entity.Product true "Product Request"
// @Success 		200 {object} entity.Product
// @Failure 		400 {object} Error
// @Failure 		404 {object} Error
// @Failure 		500 {object} Error
// @Router 			/product/{id} [put]
// @Security  ApiKeyAuth
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := pkgEntity.ParseID(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse := Error{
			Message: "Invalid request body",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
	product.ID = id

	err = h.repository.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errResponse := Error{
			Message: "Product not found",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
}

// Delete Product godoc
// @Summary 		Delete product
// @Description Delete product a product by id
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			id	path string	true	"Product ID" format(uuid)
// @Success 		200
// @Failure 		404 {object} Error
// @Failure 		500 {object} Error
// @Router 			/product/{id} [delete]
// @Security  ApiKeyAuth
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.repository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errResponse := Error{
			Message: "Product not found",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
}
