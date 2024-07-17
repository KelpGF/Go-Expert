package database

import (
	"database/sql"

	"github.com/google/uuid"
)

// ActiveRecord pattern
type Category struct {
	db *sql.DB

	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name, description string) (*Category, error) {
	id := uuid.New().String()

	stmt, err := c.db.Prepare("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description)
	if err != nil {
		return nil, err
	}

	category := Category{
		ID:          id,
		Name:        name,
		Description: description,
	}

	return &category, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}

	for rows.Next() {
		category := Category{}

		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (*Category, error) {
	stmt, err := c.db.Prepare("SELECT c.id, c.name, c.description FROM categories c INNER JOIN courses cs ON c.id = cs.category_id WHERE cs.id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	category := Category{}

	err = stmt.QueryRow(courseID).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}

	return &category, nil
}
