package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KelpGF/Go-Expert/12-GraphQL/graph"
	"github.com/KelpGF/Go-Expert/12-GraphQL/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS categories (id TEXT PRIMARY KEY, name TEXT, description TEXT)`)
	db.Exec(`CREATE TABLE IF NOT EXISTS courses (id TEXT PRIMARY KEY, title TEXT, description TEXT, category_id TEXT)`)

	courseDb := database.NewCourse(db)
	categoryDb := database.NewCategory(db)

	resolver := &graph.Resolver{
		CategoryDB: categoryDb,
		CourseDB:   courseDb,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
