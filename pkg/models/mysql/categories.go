package mysql

import (
	"database/sql"
	"errors"
	"github.com/d3z41k/shutters-calculator/pkg/models"
)

type CategoryModel struct {
	DB *sql.DB
}

func (m *CategoryModel) Insert(name string, price int) (int, error) {
	stmt := `INSERT INTO categories (name, price) 
	VALUES(?, ?))`

	result, err := m.DB.Exec(stmt, name, price)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *CategoryModel) Get(id int) (*models.Categories, error) {
	stmt := `SELECT id, name, price FROM categories WHERE id = ?`

	c := &models.Categories{}

	err := m.DB.QueryRow(stmt, id).Scan(&c.ID, &c.Name, &c.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return c, nil
}

func (m *CategoryModel) All() ([]*models.Categories, error) {
	stmt := `SELECT id, name, price FROM categories`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []*models.Categories{}

	for rows.Next() {
		c := &models.Categories{}

		err = rows.Scan(&c.ID, &c.Name, &c.Price)
		if err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
