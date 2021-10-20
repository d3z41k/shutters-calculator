package mysql

import (
	"github.com/d3z41k/shutters-calculator/pkg/models"
	"gorm.io/gorm"
)

type FabricsModel struct {
	DB *gorm.DB
}

func (m *FabricsModel) Insert(name string, categoryID int) (int, error) {
	fabric := models.Fabrics{Name: name, CategoryID: categoryID}
	result := m.DB.Create(&fabric)
	if result.Error != nil {
		return 0, result.Error
	}

	return fabric.ID, nil
}

func (m *FabricsModel) Get(id int) (*models.Fabrics, error) {
	fabric := &models.Fabrics{}
	result := m.DB.Preload("Category").First(fabric, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return fabric, nil
}

func (m *FabricsModel) All() ([]*models.Fabrics, error) {
	fabrics := []*models.Fabrics{}
	result := m.DB.Preload("Category").Find(&fabrics)
	if result.Error != nil {
		return nil, result.Error
	}

	return fabrics, nil
}
