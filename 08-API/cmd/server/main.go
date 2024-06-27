package main

import (
	"net/http"

	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"
	dbRepository "github.com/KelpGF/Go-Expert/08-APIs/internal/infrastructure/database/repository"
	"github.com/KelpGF/Go-Expert/08-APIs/internal/infrastructure/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// configs := configs.LoadConfig(".")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	mapperProductRoutes(router, db)
	mapperUserRoutes(router, db)

	http.ListenAndServe(":3000", router)
}

func mapperProductRoutes(router *chi.Mux, db *gorm.DB) {
	productRepository := dbRepository.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepository)

	router.Get("/product/{id}", productHandler.Get)
	router.Get("/product", productHandler.GetByPagination)
	router.Post("/product", productHandler.Create)
	router.Put("/product/{id}", productHandler.Update)
	router.Delete("/product/{id}", productHandler.Delete)
}

func mapperUserRoutes(router *chi.Mux, db *gorm.DB) {
	userRepository := dbRepository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepository)

	router.Post("/user", userHandler.Create)
}
