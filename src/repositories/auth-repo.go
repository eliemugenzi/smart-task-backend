package repositories

import (
	"smart-task-backend/src/db/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(user models.User) (*gorm.DB, models.User)
	FindByEmail(email string) (*gorm.DB, models.User)
	FindById(id uint) (*gorm.DB, models.User)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *authRepo {
	return &authRepo{db}
}

func (this_ *authRepo) Register(user models.User) (*gorm.DB, models.User) {
	user.Password = hashSalt(user.Password)
	userResult := this_.db.Create(&user)
	return userResult, user
}

func (this_ *authRepo) FindByEmail(email string) (*gorm.DB, models.User) {
	var user models.User
	userResult := this_.db.Where("email = ?", email).First(&user)
	return userResult, user
}

func (this_ *authRepo) FindById(id uint) (*gorm.DB, models.User) {
	var user models.User
	userResult := this_.db.First(&user, id)
	return userResult, user
}

func hashSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic("Failed to hash the password")
	}
	return string(hash)
}
