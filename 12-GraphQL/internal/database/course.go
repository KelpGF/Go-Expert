package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db *sql.DB

	ID          string
	Title       string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(title, description, categoryID string) (*Course, error) {
	id := uuid.New().String()

	stmt, err := c.db.Prepare("INSERT INTO courses (id, title, description, category_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, title, description, categoryID)
	if err != nil {
		return nil, err
	}

	course := Course{
		ID:          id,
		Title:       title,
		Description: description,
		CategoryID:  categoryID,
	}

	return &course, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, title, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []Course{}

	for rows.Next() {
		course := Course{}

		err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, title, description, category_id FROM courses WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []Course{}

	for rows.Next() {
		course := Course{}

		err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}
