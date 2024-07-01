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

// GetJwt godoc
// @Summary 		Get a JWT token
// @Description Get a JWT token by providing email and password
// @Tags 				users
// @Accept 			json
// @Produce 		json
// @Param 			request	body dto.GetJWTInput true "User Credentials"
// @Success 		200 {object} dto.GetJWTOutput
// @Failure 		400 {object} Error
// @Failure 		401 {object} Error
// @Router 			/user/generate_token [post]
func (h *UserHandlers) GetJwt(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwtAuth").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var input dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse := Error{
			Message: "Invalid request body",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	user, err := h.repository.FindByEmail(input.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		errResponse := Error{
			Message: "Invalid email or password",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if !user.ComparePassword(input.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		errResponse := Error{
			Message: "Invalid email or password",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	jwtPayload := map[string]interface{}{
		"sub":   user.ID.String(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	}
	_, token, _ := jwt.Encode(jwtPayload)

	accessToken := dto.GetJWTOutput{
		AccessToken: token,
	}

	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)
}

// Create User godoc
// @Summary 		Create a new user
// @Description Create a new user
// @Tags 				users
// @Accept 			json
// @Produce 		json
// @Param 			request	body dto.CreateUserInput true "User Request"
// @Success 		201 {object} entity.User
// @Failure 		400 {object} Error
// @Failure 		500 {object} Error
// @Router 			/user [post]
func (h *UserHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse := Error{
			Message: "Invalid request body",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	user := entity.NewUser(input.Name, input.Email, input.Password)

	err = h.repository.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := Error{
			Message: "Internal server error",
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
}
