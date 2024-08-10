package main

import (
	"context"
	"database/sql"

	"github.com/KelpGF/Go-Expert/016-Database/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(go_db_mysql:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Go",
		Description: sql.NullString{String: "Go programming language", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	loopCategories(ctx, queries)

	id := uuid.New().String()
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          id,
		Name:        "Node",
		Description: sql.NullString{String: "Node programming language", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	category, err := queries.GetCategoryById(ctx, id)
	if err != nil {
		panic(err)
	}

	println(category.ID, category.Name, category.Description.String)

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          id,
		Name:        "Node.js",
		Description: sql.NullString{String: "Node.js programming language", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	loopCategories(ctx, queries)

	err = queries.DeleteCategory(ctx, id)
	if err != nil {
		panic(err)
	}

	loopCategories(ctx, queries)
}

func loopCategories(ctx context.Context, queries *db.Queries) {
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
