package mysql

import (
	"github.com/d3z41k/shutters-calculator/pkg/models"
	"gorm.io/gorm"
)

type ProfileWidthModel struct {
	DB *gorm.DB
}

func (m *ProfileWidthModel) Insert(width, price int) (int, error) {
	profileWidth := models.ProfilesWidth{Width: width, Price: price}
	result := m.DB.Create(&profileWidth)
	if result.Error != nil {
		return 0, result.Error
	}

	return profileWidth.ID, nil
}

func (m *ProfileWidthModel) Get(id int) (*models.ProfilesWidth, error) {
	profileWidth := &models.ProfilesWidth{}
	result := m.DB.First(profileWidth, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return profileWidth, nil
}

func (m *ProfileWidthModel) All() ([]*models.ProfilesWidth, error) {
	var profileWidths []*models.ProfilesWidth
	result := m.DB.Find(&profileWidths)
	if result.Error != nil {
		return nil, result.Error
	}

	return profileWidths, nil
}
