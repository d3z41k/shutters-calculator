package mysql

import (
	"github.com/d3z41k/shutters-calculator/pkg/models"
	"gorm.io/gorm"
)

type ProfilesModel struct {
	DB *gorm.DB
}

func (m *ProfilesModel) Insert(width, price int) (int, error) {
	profile := models.Profiles{Width: width, Price: price}
	result := m.DB.Create(&profile)
	if result.Error != nil {
		return 0, result.Error
	}

	return profile.ID, nil
}

func (m *ProfilesModel) Get(id int) (*models.Profiles, error) {
	profile := &models.Profiles{}
	result := m.DB.First(profile, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return profile, nil
}

func (m *ProfilesModel) GetByFilter(filter map[string]interface{}) (*models.Profiles, error) {
	profile := &models.Profiles{}
	result := m.DB.Where(filter).First(profile)
	if result.Error != nil {
		return nil, result.Error
	}

	return profile, nil
}

func (m *ProfilesModel) All() ([]*models.Profiles, error) {
	var profiles []*models.Profiles
	result := m.DB.Find(&profiles)
	if result.Error != nil {
		return nil, result.Error
	}

	return profiles, nil
}
