package main

import (
	"log"
	"net/http"

	"github.com/KelpGF/Go-Expert/08-APIs/configs"
	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"
	dbRepository "github.com/KelpGF/Go-Expert/08-APIs/internal/infrastructure/database/repository"
	"github.com/KelpGF/Go-Expert/08-APIs/internal/infrastructure/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/KelpGF/Go-Expert/08-APIs/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Go Expert API
// @version         1.0
// @description     Product API with authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Kelvin Gomes
// @contact.url    https://www.linkedin.com/in/kelvin-gomes-fernandes
// @contact.email  kelvingomesdeveloper@gmail.com

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs := configs.LoadConfig(".")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(LogRequest)
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwtAuth", configs.TokenAuth))
	router.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))

	mapperProductRoutes(router, db, configs.TokenAuth)
	mapperUserRoutes(router, db)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:3000/docs/doc.json")))

	http.ListenAndServe(configs.WebServerHost+":"+configs.WebServerPort, router)
}

func mapperProductRoutes(router *chi.Mux, db *gorm.DB, jwt *jwtauth.JWTAuth) {
	productRepository := dbRepository.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepository)

	router.Route("/product", func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwt))
		r.Use(jwtauth.Authenticator)

		r.Get("/{id}", productHandler.Get)
		r.Get("/", productHandler.GetByPagination)
		r.Post("/", productHandler.Create)
		r.Put("/{id}", productHandler.Update)
		r.Delete("/{id}", productHandler.Delete)
	})
}

func mapperUserRoutes(router *chi.Mux, db *gorm.DB) {
	userRepository := dbRepository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepository)

	router.Post("/user", userHandler.Create)
	router.Post("/user/generate_token", userHandler.GetJwt)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request")
		next.ServeHTTP(w, r)
	})
}
