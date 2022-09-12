package repositories

import (
	"dumflix/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	// Login(user models.User) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

// func (r *repository) Login(email string) (models.User, error) {
// 	var user models.User
// 	err := r.db.First(&user, "email=?", email).Error

// 	return user, err
// }
