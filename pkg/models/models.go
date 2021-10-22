package models

import "errors"

var (
	ErrNoRecord = errors.New("models: no matching record found")
	//ErrInvalidCredentials = errors.New("models: invalid credentials")
	//ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Tabler interface {
	TableName() string
}

type Categories struct {
	ID    int
	Name  string
	Price int
}

type Fabrics struct {
	ID         int
	Name       string
	CategoryID int
	Category   Categories
}

func (ProfilesWidth) TableName() string {
	return "profiles_width"
}

type ProfilesWidth struct {
	ID    int
	Width int
	Price int
}

func (ProfilesHeight) TableName() string {
	return "profiles_height"
}

type ProfilesHeight struct {
	ID     int
	Height int
}
