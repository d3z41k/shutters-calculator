package models

import "errors"

var (
	ErrNoRecord = errors.New("models: no matching record found")
	//ErrInvalidCredentials = errors.New("models: invalid credentials")
	//ErrDuplicateEmail     = errors.New("models: duplicate email")
)

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
