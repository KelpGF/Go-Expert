package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/dto"
	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"
	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/repository"
)

type UserHandlers struct {
	repository repository.UserRepository
}

func NewUserHandler(repository repository.UserRepository) *UserHandlers {
	return &UserHandlers{repository: repository}
}

func (h *UserHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := entity.NewUser(input.Name, input.Email, input.Password)

	err = h.repository.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
}
