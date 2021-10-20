package mysql

import (
	"github.com/d3z41k/shutters-calculator/pkg/models"
	"gorm.io/gorm"
)

type CategoryModel struct {
	DB *gorm.DB
}

func (m *CategoryModel) Insert(name string, price int) (int, error) {
	category := models.Categories{Name: name, Price: price}
	result := m.DB.Create(&category)
	if result.Error != nil {
		return 0, result.Error
	}

	return category.ID, nil
}

func (m *CategoryModel) Get(id int) (*models.Categories, error) {
	category := &models.Categories{}
	result := m.DB.First(category, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (m *CategoryModel) All() ([]*models.Categories, error) {
	categories := []*models.Categories{}
	result := m.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}
