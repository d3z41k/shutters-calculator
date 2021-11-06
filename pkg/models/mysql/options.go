package mysql

import (
	"github.com/d3z41k/shutters-calculator/pkg/models"
	"gorm.io/gorm"
)

type OptionsModel struct {
	DB *gorm.DB
}

func (m *OptionsModel) Insert(key string, price int) (int, error) {
	option := models.Options{Key: key, Price: price}
	result := m.DB.Create(&option)
	if result.Error != nil {
		return 0, result.Error
	}

	return option.ID, nil
}

func (m *OptionsModel) Get(id int) (*models.Options, error) {
	option := &models.Options{}
	result := m.DB.First(option, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return option, nil
}

func (m *OptionsModel) GetSumByKeys(keys []string) (int, error) {
	type optionsSum struct {
		Total int
	}

	var os optionsSum

	result := m.DB.Model(&models.Options{}).Select("sum(price) as total").Where("`key` IN ?", keys).Find(&os)
	if result.Error != nil {
		return 0, result.Error
	}

	return os.Total, nil
}

func (m *OptionsModel) All() ([]*models.Options, error) {
	var options []*models.Options
	result := m.DB.Find(&options)
	if result.Error != nil {
		return nil, result.Error
	}

	return options, nil
}
