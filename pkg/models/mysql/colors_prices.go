package mysql

import (
	"github.com/d3z41k/shutters-calculator/pkg/models"
	"gorm.io/gorm"
)

type ColorsPricesModel struct {
	DB *gorm.DB
}

func (m *ColorsPricesModel) Insert(colorId, width, price int) (int, error) {
	colorPrice := models.ColorsPrices{ColorID: colorId, Width: width, Price: price}
	result := m.DB.Create(&colorPrice)
	if result.Error != nil {
		return 0, result.Error
	}

	return colorPrice.ID, nil
}

func (m *ColorsPricesModel) Get(id int) (*models.ColorsPrices, error) {
	colorPrice := &models.ColorsPrices{}
	result := m.DB.First(colorPrice, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return colorPrice, nil
}

func (m *ColorsPricesModel) GetByFilter(filter map[string]interface{}) (*models.ColorsPrices, error) {
	colorPrice := &models.ColorsPrices{}
	result := m.DB.Where(filter).First(colorPrice)

	if result.Error != nil {
		return nil, result.Error
	}

	return colorPrice, nil
}

func (m *ColorsPricesModel) All() ([]*models.ColorsPrices, error) {
	var colorsPrices []*models.ColorsPrices
	result := m.DB.Find(&colorsPrices)
	if result.Error != nil {
		return nil, result.Error
	}

	return colorsPrices, nil
}
