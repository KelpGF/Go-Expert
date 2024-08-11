package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/KelpGF/Go-Expert/016-Database/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourse(ctx context.Context, category CategoryParams, course CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error

		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryID:  category.ID,
			Price:       course.Price,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(go_db_mysql:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	courses, err := queries.ListCourses(ctx)
	for _, course := range courses {
		fmt.Println(course.ID, course.Name, course.Description.String, course.Price, course.CategoryName)
	}

	// courseParams := CourseParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Go",
	// 	Description: sql.NullString{String: "Go programming language", Valid: true},
	// 	Price:       100.00,
	// }

	// categoryParams := CategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "GoLang",
	// 	Description: sql.NullString{String: "List of Go courses", Valid: true},
	// }

	// courseDB := NewCourseDB(dbConn)

	// err = courseDB.CreateCourse(ctx, categoryParams, courseParams)
	// if err != nil {
	// 	panic(err)
	// }
}
