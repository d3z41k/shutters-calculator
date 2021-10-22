package mysql

import (
	"github.com/d3z41k/shutters-calculator/pkg/models"
	"gorm.io/gorm"
)

type ProfileHeightModel struct {
	DB *gorm.DB
}

func (m *ProfileHeightModel) Insert(height int) (int, error) {
	profileHeight := models.ProfilesHeight{Height: height}
	result := m.DB.Create(&profileHeight)
	if result.Error != nil {
		return 0, result.Error
	}

	return profileHeight.ID, nil
}

func (m *ProfileHeightModel) Get(id int) (*models.ProfilesHeight, error) {
	profileHeight := &models.ProfilesHeight{}
	result := m.DB.First(profileHeight, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return profileHeight, nil
}

func (m *ProfileHeightModel) All() ([]*models.ProfilesHeight, error) {
	var profileHeights []*models.ProfilesHeight
	result := m.DB.Find(&profileHeights)
	if result.Error != nil {
		return nil, result.Error
	}

	return profileHeights, nil
}
