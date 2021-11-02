package mysql

import (
	"github.com/d3z41k/shutters-calculator/pkg/models"
	"gorm.io/gorm"
)

type ColorsModel struct {
	DB *gorm.DB
}

func (m *ColorsModel) Insert(name string) (int, error) {
	color := models.Colors{Name: name}
	result := m.DB.Create(&color)
	if result.Error != nil {
		return 0, result.Error
	}

	return color.ID, nil
}

func (m *ColorsModel) Get(id int) (*models.Colors, error) {
	color := &models.Colors{}
	result := m.DB.First(color, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return color, nil
}

func (m *ColorsModel) All() ([]*models.Colors, error) {
	var colors []*models.Colors
	result := m.DB.Find(&colors)
	if result.Error != nil {
		return nil, result.Error
	}

	return colors, nil
}
