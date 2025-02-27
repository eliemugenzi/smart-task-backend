package services

import (
	"smart-task-backend/src/db/models"
	"smart-task-backend/src/dto"
	"smart-task-backend/src/repositories"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(userDto dto.User) (*gorm.DB, models.User)
	VerifyCredential(loginDto dto.Login) (bool, uint64)
	FindUserByEmail(email string) models.User
	GetUsers() []models.User
}

type authService struct {
	authRepo repositories.AuthRepo
}

func NewAuthService(authRepo repositories.AuthRepo) *authService {
	return &authService{
		authRepo,
	}
}

func (this_ *authService) Register(userDto dto.User) (*gorm.DB, models.User) {
	userModel := models.User{}
	err := smapping.FillStruct(&userModel, smapping.MapFields(&userDto))
	if err != nil {
		panic(err)
	}
	return this_.authRepo.Register(userModel)
}

func comparePassword(hashedPass []byte, plainPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPass, plainPass)

	return err == nil
}

func (this_ *authService) FindUserByEmail(email string) models.User {
	_, user := this_.authRepo.FindByEmail(email)
	return user
}

func (this_ *authService) VerifyCredential(loginDto dto.Login) (bool, uint64) {
	result, user := this_.authRepo.FindByEmail(loginDto.Email)
	if result.Error == nil && user.ID != 0 {
		return comparePassword([]byte(user.Password), []byte(loginDto.Password)), uint64(user.ID)
	}
	return false, 0
}

func (this_ *authService) GetUsers() []models.User {
	_, users := this_.authRepo.GetUsers()

	return users
}
