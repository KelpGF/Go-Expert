package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/dto"
	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"
	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/repository"
	"github.com/go-chi/jwtauth"
)

type UserHandlers struct {
	repository   repository.UserRepository
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(repository repository.UserRepository) *UserHandlers {
	return &UserHandlers{
		repository: repository,
	}
}

func (h *UserHandlers) GetJwt(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwtAuth").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var input dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.repository.FindByEmail(input.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.ComparePassword(input.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwtPayload := map[string]interface{}{
		"sub":   user.ID.String(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	}
	_, token, _ := jwt.Encode(jwtPayload)

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)
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
